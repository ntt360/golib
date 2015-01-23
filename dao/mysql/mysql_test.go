/**
/ @file mysql_test.go
* @brief test mysql
* @author ligang
* @version 1.0
* @date 2015-01-05
*/

package mysql

import (
	"github.com/mydoraemon/golib/bizlog"
	"testing"
)

func init() {
	bizlog.Init("/home/ligang/devspace/golib/logs")
	bizlog.NewLogger("mysql", bizlog.MODE_SYNC, "test", bizlog.SPLIT_BY_DAY, 0, bizlog.MSG_FMT_LINE_HEADER)
}

func TestUpdateById(t *testing.T) {
}
