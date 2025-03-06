package counts

import (
	"bufio"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"

	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
)

func CountRegMatchs(logDir string, re *regexp.Regexp) int {
	counts := []int{}
	err := filepath.Walk(logDir, func(path string, d fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		// read file
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		count, err := match(file, re)
		if err != nil {
			return err
		}
		logrus.WithField("path", path).Infof("count: %d", count)
		counts = append(counts, count)
		return nil
	})
	if err != nil {
		panic(err)
	}
	return lo.Sum(counts)
}

func match(r io.Reader, re *regexp.Regexp) (int, error) {
	count := 0
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		// 正则匹配
		matches := re.FindStringSubmatch(line)
		if len(matches) > 0 {
			count++
		}
	}
	if err := scanner.Err(); err != nil {
		return 0, errors.WithMessage(err, "读取文件错误")
	}
	return count, nil
}
