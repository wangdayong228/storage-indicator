package extractors

import (
	"bufio"
	"encoding/csv"
	"io"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/nft-rainbow/rainbow-goutils/utils/commonutils"
	"github.com/pkg/errors"
	"github.com/samber/lo"
)

// 同步进度差异
func SyncProgressDiff(r io.Reader, w csv.Writer) error {
	w.Write([]string{"Timestamp", "FromBlockNumber", "LatestBlockNumber", "Diff"})

	re := regexp.MustCompile(`^(\S+Z).*?from block number (\d+), latest block number (\d+)`)
	// 逐行读取日志文件
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		// 正则匹配
		matches := re.FindStringSubmatch(line)
		if len(matches) == 4 {
			// 写入 CSV
			diff := commonutils.Must(strconv.Atoi(matches[3])) - commonutils.Must(strconv.Atoi(matches[2]))
			if err := w.Write([]string{matches[1], matches[2], matches[3], strconv.Itoa(diff)}); err != nil {
				return err
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return errors.WithMessage(err, "读取文件错误")
	}

	return nil
}

type MemPoolRefreshPeroid struct {
	TxSeq     string
	StartTime string
	EndTime   string
}

// 内存池刷新效率
func MemPoolRefreshRate(r io.Reader, w csv.Writer) error {
	w.Write([]string{"TxSeq", "StratTime", "EndTime", "TimeUse(us)"})

	peroids := make(map[string]*MemPoolRefreshPeroid)

	startRe := regexp.MustCompile(`^(\S+Z).*start to flush cached segments to log store.*tx_seq:(\d+)`)
	endRe := regexp.MustCompile(`^(\S+Z).*?cached segments flushed to log store.*?tx_seq:(\d+)`)
	// 逐行读取日志文件
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		// 正则匹配
		matches := startRe.FindStringSubmatch(line)

		if len(matches) == 3 {
			startTime, txSeq := matches[1], matches[2]
			if _, ok := peroids[txSeq]; !ok {
				peroids[txSeq] = &MemPoolRefreshPeroid{}
			}
			peroids[txSeq].TxSeq = txSeq
			peroids[txSeq].StartTime = startTime
		}

		matches = endRe.FindStringSubmatch(line)
		if len(matches) == 3 {
			endTime, txSeq := matches[1], matches[2]
			if _, ok := peroids[txSeq]; !ok {
				peroids[txSeq] = &MemPoolRefreshPeroid{}
			}
			peroids[txSeq].EndTime = endTime
		}
	}

	if err := scanner.Err(); err != nil {
		return errors.WithMessage(err, "读取文件错误")
	}

	txSeqs := lo.Keys(peroids)
	slices.Sort(txSeqs)

	for _, txSeq := range txSeqs {
		peroid := peroids[txSeq]
		timeUse := commonutils.Must(time.Parse(time.RFC3339, peroid.EndTime)).Sub(commonutils.Must(time.Parse(time.RFC3339, peroid.StartTime)))
		if err := w.Write([]string{peroid.TxSeq, peroid.StartTime, peroid.EndTime, strconv.Itoa(int(timeUse.Microseconds()))}); err != nil {
			return err
		}
	}

	return nil
}

type TxSyncCompletePeroid struct {
	TxSeq     string
	StartTime string
	EndTime   string
}

// 事务同步成功的同步时间
func TxSyncCompleteTimeCost(r io.Reader, w csv.Writer) error {
	w.Write([]string{"TxSeq", "StratTime", "EndTime", "TimeUse(sec)"})

	peroids := make(map[string]*TxSyncCompletePeroid)

	// 2025-03-02T02:39:43.690843Z  INFO sync::service: Start to sync file tx_seq=3870740 maybe_range=None maybe_peer=None
	startRe := regexp.MustCompile(`^(\S+Z).*?Start to sync file tx_seq=(\d+) maybe_range=None maybe_peer=None`)
	// 2025-03-02T00:00:00.953788Z DEBUG sync::auto_sync::batcher_random: Completed to sync file, state = Ok(RandomBatcherState { name: "random_historical", tasks: [5253251, 5253256, 5253258, 5253259, 5253260, 5253261, 5253262], pending_txs: 2836315, ready_txs: 10, cached_ready_txs: 0 }) tx_seq=5253257 sync_result=Completed
	endRe := regexp.MustCompile(`^(\S+Z).*?Completed to sync file.* tx_seq=(\d+) sync_result=Completed`)
	// 逐行读取日志文件
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		// 正则匹配
		matches := startRe.FindStringSubmatch(line)

		if len(matches) == 3 {
			startTime, txSeq := matches[1], matches[2]
			if _, ok := peroids[txSeq]; !ok {
				peroids[txSeq] = &TxSyncCompletePeroid{}
			}
			peroids[txSeq].TxSeq = txSeq
			peroids[txSeq].StartTime = startTime

		}
		matches = endRe.FindStringSubmatch(line)
		if len(matches) == 3 {
			endTime, txSeq := matches[1], matches[2]
			if _, ok := peroids[txSeq]; !ok {
				peroids[txSeq] = &TxSyncCompletePeroid{}
			}
			peroids[txSeq].EndTime = endTime
		}
	}

	if err := scanner.Err(); err != nil {
		return errors.WithMessage(err, "读取文件错误")
	}

	txSeq := lo.Keys(peroids)
	slices.Sort(txSeq)

	for _, txSeq := range txSeq {
		peroid := peroids[txSeq]
		if peroid.StartTime == "" || peroid.EndTime == "" {
			continue
		}
		timeUse := commonutils.Must(time.Parse(time.RFC3339, peroid.EndTime)).Sub(commonutils.Must(time.Parse(time.RFC3339, peroid.StartTime)))
		if err := w.Write([]string{peroid.TxSeq, peroid.StartTime, peroid.EndTime, strconv.Itoa(int(timeUse.Seconds()))}); err != nil {
			return err
		}
	}

	return nil
}

// 同步任务队列积压
func SyncTaskBacklog(r io.Reader, w csv.Writer) error {
	w.Write([]string{"Timestamp", "Incompleted_count"})

	// 2025-03-02T05:10:57.343045Z DEBUG sync::service: Sync stat: incompleted = [3702335, 3678838, 5266919, 3018933, 3963072, 3383036, 5266914, 3254619, 5266924, 3485681, 2867380], completed = []
	re := regexp.MustCompile(`^(\S+Z).*?Sync stat: incompleted = \[(.*)\], completed =.*`)
	// 逐行读取日志文件
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		// 正则匹配
		matches := re.FindStringSubmatch(line)
		if len(matches) == 3 {
			// 写入 CSV
			len := len(strings.Split(matches[2], ","))
			if err := w.Write([]string{matches[1], strconv.Itoa(len)}); err != nil {
				return err
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return errors.WithMessage(err, "读取文件错误")
	}

	return nil
}

type MineWorkInfo struct {
	Timestamp      string
	ScratchPad     string
	ScratchPadRate float32
	Loading        string
	LoadingRate    float32
	PadMix         string
	PadMixRate     float32
	Hit            string
	HitRate        float32
}

// 挖矿阶段耗时分布
func MineWork(r io.Reader, w csv.Writer) error {
	w.Write([]string{"Timestamp", "ScratchPad(us)", "ScratchPadRate", "Loading(us)", "LoadingRate", "PadMix(us)", "PadMixRate", "Hit(us)", "HitRate"})

	re := regexp.MustCompile(`^(\S+Z).*?Mine iterations statistics: scratch pad: (\d+), loading: (\d+), pad_mix: (\d+), hit: (\d+)`)

	infos := []MineWorkInfo{}
	// 逐行读取日志文件
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		// 正则匹配
		matches := re.FindStringSubmatch(line)
		if len(matches) == 6 {
			// 写入 CSV
			infos = append(infos, MineWorkInfo{
				Timestamp:  matches[1],
				ScratchPad: matches[2],
				Loading:    matches[3],
				PadMix:     matches[4],
				Hit:        matches[5],
			})
		}
	}

	for i, info := range infos {
		if i == 0 {
			continue
		}

		timeUse := commonutils.Must(time.Parse(time.RFC3339, info.Timestamp)).Sub(commonutils.Must(time.Parse(time.RFC3339, infos[i-1].Timestamp)))

		scratchPadDiff := commonutils.Must(strconv.Atoi(info.ScratchPad)) - commonutils.Must(strconv.Atoi(infos[i-1].ScratchPad))
		info.ScratchPadRate = float32(scratchPadDiff) / float32(timeUse.Microseconds()) * 1e6
		loadingDiff := commonutils.Must(strconv.Atoi(info.Loading)) - commonutils.Must(strconv.Atoi(infos[i-1].Loading))
		info.LoadingRate = float32(loadingDiff) / float32(timeUse.Microseconds()) * 1e6
		padMixDiff := commonutils.Must(strconv.Atoi(info.PadMix)) - commonutils.Must(strconv.Atoi(infos[i-1].PadMix))
		info.PadMixRate = float32(padMixDiff) / float32(timeUse.Microseconds()) * 1e6
		hitDiff := commonutils.Must(strconv.Atoi(info.Hit)) - commonutils.Must(strconv.Atoi(infos[i-1].Hit))
		info.HitRate = float32(hitDiff) / float32(timeUse.Microseconds()) * 1e6

		if err := w.Write([]string{info.Timestamp,
			info.ScratchPad,
			strconv.FormatFloat(float64(info.ScratchPadRate), 'f', 2, 64),
			info.Loading,
			strconv.FormatFloat(float64(info.LoadingRate), 'f', 2, 64),
			info.PadMix,
			strconv.FormatFloat(float64(info.PadMixRate), 'f', 2, 64),
			info.Hit,
			strconv.FormatFloat(float64(info.HitRate), 'f', 2, 64),
		}); err != nil {
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		return errors.WithMessage(err, "读取文件错误")
	}

	return nil
}
