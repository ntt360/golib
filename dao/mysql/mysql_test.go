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
	bizlog.NewLogger("test", bizlog.MODE_SYNC, "", bizlog.SPLIT_BY_DAY, 0, bizlog.MSG_FMT_LINE_HEADER)
}

func TestUpdateById(t *testing.T) {
	mysql_conf := executor.T_Mysql_Conf{
		Host: "127.0.0.1",
		User: "root",
		Pass: "123",
		Port: "3306",
		Name: "test",
	}

	dao := NewBaseDao("test", mysql_conf, bizlog.GetLogger("test"))
	dao.UpdateById("test", 8, []string{"id"}, 8)
}
