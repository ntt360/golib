/**
* @file mysql_test.go
* @brief test mysql executor
* @author ligang
* @version 1.0
* @date 2014-12-22
 */

package mysql

import (
	"fmt"
	"github.com/mydoraemon/golib/bizlog"
	"testing"
)

func init() {
	bizlog.Init("/home/ligang/devspace/golib/logs")
	bizlog.NewLogger("test", bizlog.MODE_SYNC, "", bizlog.SPLIT_BY_DAY, 0, bizlog.MSG_FMT_LINE_HEADER)
}

func TestQuery(t *testing.T) {
	executor := getExecutor()
	sql := "select * from test where id = ?"

	rows, err := executor.Query(sql, 8)
	fmt.Println(err)
	for rows.Next() {
		var id int
		rows.Scan(&id)
		fmt.Println(id)
	}
}

func TestExec(t *testing.T) {
	executor := getExecutor()
	sql := "update test set id = ? where id = ?"

	result, err := executor.Exec(sql, 8, 8)
	fmt.Println(err)
	fmt.Println(result.RowsAffected())
}

func getExecutor() *T_Mysql_Executor {
	mysql_conf := T_Mysql_Conf{
		Host: "127.0.0.1",
		User: "root",
		Pass: "123",
		Port: "3306",
		Name: "test",
	}

	return NewMysqlExecutor("test", mysql_conf, bizlog.GetLogger("test"))
}
