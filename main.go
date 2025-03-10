package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/wangdayong228/storage-indicator/extractors"
)

func main() {

	// 全天
	source := "./log/zgs.log.2025-03-02"
	outDir := "./out"

	// 少量日志
	// source := "./log/zgs.log.2025-03-02.short"
	// outDir := "./out/short"
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

	fmt.Println("CSV 生成完成：output.csv")
	return nil
}
