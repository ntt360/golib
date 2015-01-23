package bizlog

import (
	"testing"
)

func init() {
	Init("/home/ligang/devspace/golib/logs")
}

func TestSyncLog(t *testing.T) {
	logger := NewLogger("test", MODE_SYNC, "", SPLIT_BY_DAY, DEF_BUFSIZE, MSG_FMT_LINE_HEADER)
	logger.Log("test sync log")
	logger.Flush()
}

func TestAsyncLog(t *testing.T) {
	logger := NewLogger("test", MODE_ASYNC, "", SPLIT_BY_DAY, DEF_BUFSIZE, MSG_FMT_LINE_HEADER)
	logger.Log("test async log")

	FlushAll()
}
