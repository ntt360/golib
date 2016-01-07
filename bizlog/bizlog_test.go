package bizlog

import (
	"testing"
)

func init() {
	Init("/tmp/bizlog")
}

func TestSyncLog(t *testing.T) {
	logger := NewLogger("test", MODE_SYNC, "", SPLIT_BY_DAY, DEF_BUFSIZE, MSG_FMT_LINE_HEADER)
	logger.Log("test sync log")
	logger.Free()
}

func TestAsyncLog(t *testing.T) {
	logger := NewLogger("test", MODE_ASYNC, "", SPLIT_BY_DAY, DEF_BUFSIZE, MSG_FMT_LINE_HEADER)
	logger.Log("test async log")

	Free()
}
