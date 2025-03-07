package counts

import (
	"regexp"
	"testing"
	"time"
)

func TestSyncProgressDiff(t *testing.T) {
	count := CountRegMatchs("SyncProgressDiff", "/Users/dayong/myspace/mywork/storage-indicator/log", time.Now(), regexp.MustCompile(`^(\S+Z).*?from block number (\d+), latest block number (\d+)`))
	t.Logf("count: %d", count)
}
