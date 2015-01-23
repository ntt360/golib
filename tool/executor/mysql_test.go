/**
* @file mysql_test.go
* @brief test mysql executor
* @author ligang
* @version 1.0
* @date 2014-12-22
 */

package executor

import (
	"fmt"
	"mydoraemon/golib/bizlog"
	"mydoraemon/golib/dao"
	"testing"
)

func init() {
	bizlog.Init("/home/ligang/devspace/golib/logs")
	bizlog.NewLogger("mysql", bizlog.MODE_SYNC, "", bizlog.SPLIT_BY_DAY, 0, bizlog.MSG_FMT_LINE_HEADER)
}

func TestQuery(t *testing.T) {
	executor := getExecutor()
	sql := "select * from sync_status where role = ?"
	values := []interface{}{2}

	rows := executor.Query(sql, values...)
	for rows.Next() {
		item := new(dao.T_Sync_Status)
		rows.Scan(&item.Id, &item.Host, &item.Idc, &item.Role, &item.Status, &item.Last_Term)
		fmt.Println(item)
	}
}

func TestExec(t *testing.T) {
	executor := getExecutor()
	sql := "update sync_status set role = ? where id = ?"
	values := []interface{}{2, 1}

	result := executor.Exec(sql, values...)
	fmt.Println(result.RowsAffected())
}

func getExecutor() *T_Mysql_Executor {
	mysql_conf := config.GetMysqlConf()
	conf := T_Mysql_Conf{
		Host: "127.0.0.1",
		User: "root",
		Pass: "123",
		Port: "3306",
		Name: "test",
	}

	return NewMysqlExecutor("test", conf, bizlog.GetLogger("test"))
}
