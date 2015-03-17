/**
* @file mysql.go
* @brief dao for mysql
* @author ligang
* @version 1.0
* @date 2014-12-29
 */

package mysql

import (
	//     "fmt"
	"errors"
	"github.com/mydoraemon/golib/bizlog"
	executor "github.com/mydoraemon/golib/executor/mysql"
	"strconv"
	"strings"
)

type T_Base_Dao struct {
	Executor *executor.T_Mysql_Executor
}

func NewBaseDao(key string, mysql_conf executor.T_Mysql_Conf, logger bizlog.I_Logger) *T_Base_Dao {
	if nil == logger {
		logger = bizlog.NewLogger(LOG_KEY_MYSQL, bizlog.MODE_ASYNC, "", bizlog.SPLIT_BY_DAY, bizlog.DEF_BUFSIZE, bizlog.MSG_FMT_LINE_HEADER)
	}

	return &T_Base_Dao{
		Executor: executor.NewMysqlExecutor(key, mysql_conf, logger),
	}
}

func (dao *T_Base_Dao) UpdateById(table string, id int, cols []string, values ...interface{}) (int64, error) {
	if 0 == len(cols) {
		return 0, errors.New("invalid cols")
	}

	sql := "update " + table + " set"
	for _, col := range cols {
		sql += " " + col + " = ?,"
	}
	sql = strings.TrimRight(sql, ",")
	sql += " where id = " + strconv.Itoa(id)

	result, err := dao.Executor.Exec(sql, values...)
	if nil != err {
		return 0, err
	}

	affected, _ := result.RowsAffected()
	return affected, nil
}
