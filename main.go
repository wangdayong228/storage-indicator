package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/wangdayong228/storage-indicator/counts"
	"github.com/wangdayong228/storage-indicator/extractors"
)

func main() {
	// 定义命令行参数
	extractFlag := flag.Bool("extract", false, "执行提取操作")
	countFlag := flag.Bool("count", false, "执行计数操作")
	sourceFlag := flag.String("source", "", "日志目录")
	outFlag := flag.String("out", "", "输出目录")
	startTimeFlag := flag.String("start", "", "开始时间")
	flag.Parse()

	logrus.WithFields(logrus.Fields{
		"extractFlag":   *extractFlag,
		"countFlag":     *countFlag,
		"sourceFlag":    *sourceFlag,
		"outFlag":       *outFlag,
		"startTimeFlag": *startTimeFlag,
	}).Infof("main")
	// 全天

	// 根据命令行参数执行不同的操作
	if *extractFlag {
		// source := "./log/zgs.log.2025-03-02"
		fmt.Println("执行提取操作")
		ExtarctIndicators(*sourceFlag, *outFlag)
	} else if *countFlag {
		// 这里可以添加 count 操作的实现

		startTime, err := time.Parse("2006-01-02", *startTimeFlag)
		if err != nil {
			logrus.WithError(err).Error("parse start time error")
			os.Exit(1)
		}
		fmt.Printf("执行计数操作，开始时间：%s\n", startTime.Format("2006-01-02"))
		CountIndicators(*sourceFlag, startTime)
	} else {
		fmt.Println("请指定一个操作: -extract 或 -count")
	}

	// 少量日志
	// source := "./log/zgs.log.2025-03-02.short"
	// outDir := "./out/short"

}

func CountIndicators(source string, startTime time.Time) {
	var (
		SyncProgressDiffCount  int
		MemPoolRefreshRate     int
		TxSyncCompleteTimeCost int
		SyncTaskBacklog        int
		MineWork               int
	)

	var wg sync.WaitGroup
	wg.Add(5)

	go func() {
		defer wg.Done()
		SyncProgressDiffCount = counts.CountRegMatchs("SyncProgressDiff", source, startTime, regexp.MustCompile(`^(\S+Z).*?from block number (\d+), latest block number (\d+)`))
	}()

	go func() {
		defer wg.Done()
		MemPoolRefreshRate = counts.CountRegMatchs("MemPoolRefreshRate", source, startTime, regexp.MustCompile(`^(\S+Z).*?cached segments flushed to log store.*?tx_seq:(\d+)`))
	}()

	go func() {
		defer wg.Done()
		TxSyncCompleteTimeCost = counts.CountRegMatchs("TxSyncCompleteTimeCost", source, startTime, regexp.MustCompile(`^(\S+Z).*?Completed to sync file.* tx_seq=(\d+) sync_result=Completed`))
	}()

	go func() {
		defer wg.Done()
		SyncTaskBacklog = counts.CountRegMatchs("SyncTaskBacklog", source, startTime, regexp.MustCompile(`^(\S+Z).*?Sync stat: incompleted = \[(.*)\], completed =.*`))
	}()

	go func() {
		defer wg.Done()
		MineWork = counts.CountRegMatchs("MineWork", source, startTime, regexp.MustCompile(`^(\S+Z).*?Mine iterations statistics: scratch pad: (\d+), loading: (\d+), pad_mix: (\d+), hit: (\d+)`))
	}()

	wg.Wait()

	logrus.WithFields(logrus.Fields{
		"SyncProgressDiffCount":  SyncProgressDiffCount,
		"MemPoolRefreshRate":     MemPoolRefreshRate,
		"TxSyncCompleteTimeCost": TxSyncCompleteTimeCost,
		"SyncTaskBacklog":        SyncTaskBacklog,
		"MineWork":               MineWork,
	}).Infof("Count indicators completed")
}

func ExtarctIndicators(source, outDir string) {
	extract(extractors.SyncProgressDiff, source, outDir+"/SyncProgressDiff.csv")
	extract(extractors.MemPoolRefreshRate, source, outDir+"/MemPoolRefreshRate.csv")
	extract(extractors.TxSyncCompleteTimeCost, source, outDir+"/TxSyncCompleteTimeCost.csv")
	extract(extractors.SyncTaskBacklog, source, outDir+"/SyncTaskBacklog.csv")
	extract(extractors.MineWork, source, outDir+"/MineWork.csv")
}

func extract(fn func(r io.Reader, w csv.Writer) error, sourceFile string, targetFile string) error {
	// 打开日志文件
	file, err := os.Open(sourceFile)
	if err != nil {
		fmt.Println("无法打开文件:", err)
		return err
	}
	defer file.Close()

	// 创建 CSV 输出文件
	csvFile, err := os.Create(targetFile)
	if err != nil {
		fmt.Println("无法创建 CSV 文件:", err)
		return err
	}
	defer csvFile.Close()
	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	if err := fn(file, *writer); err != nil {
		fmt.Println("提取错误:", err)
		return err
	}

	fmt.Printf("CSV 生成完成：%s.csv\n", targetFile)
	return nil
}
