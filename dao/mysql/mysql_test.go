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
	"github.com/mydoraemon/golib/tool/executor"
	"testing"
)

func init() {
	bizlog.Init("/home/ligang/devspace/golib/logs")
	bizlog.NewLogger(LOG_KEY_MYSQL, bizlog.MODE_SYNC, "", bizlog.SPLIT_BY_DAY, 0, bizlog.MSG_FMT_LINE_HEADER)
}

func TestUpdateById(t *testing.T) {
	key = "test"
	mysql_conf := T_Mysql_Conf{
		Host: "127.0.0.1",
		User: "root",
		Pass: "123",
		Port: "3306",
		Name: "test",
	}
}
