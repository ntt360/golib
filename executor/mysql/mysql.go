/**
* @file mysql.go
* @brief mysql executor
* @author ligang
* @version 1.0
* @date 2014-12-22
 */

package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mydoraemon/golib/bizlog"
	"strings"
)

type T_Mysql_Conf struct {
	Host string
	User string
	Pass string
	Port string
	Name string
}

type T_Mysql_Executor struct {
	key    string
	logger bizlog.I_Logger
	db     *sql.DB
}

type I_Mysql_Row interface {
	Scan(dest ...interface{}) error
}

type I_Mysql_Rows interface {
	Scan(dest ...interface{}) error
	Next() bool
}

type I_Mysql_Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

var _mysql_executor_container map[string]*T_Mysql_Executor

func init() {
	_mysql_executor_container = make(map[string]*T_Mysql_Executor)
}

/**
* @brief one db one executor
*
* @param string
* @param string
* @param string
* @param string
* @param string
* @param string
*
* @return
 */
func NewMysqlExecutor(key string, mysql_conf T_Mysql_Conf, logger bizlog.I_Logger) *T_Mysql_Executor {
	if nil != _mysql_executor_container[key] {
		return GetMysqlExecutor(key)
	}

	mysql_executor := new(T_Mysql_Executor)
	mysql_executor.key = key
	mysql_executor.logger = logger

	dsn := mysql_conf.User + ":" + mysql_conf.Pass + "@tcp(" + mysql_conf.Host + ":" + mysql_conf.Port + ")/" + mysql_conf.Name
	mysql_executor.db, _ = sql.Open("mysql", dsn)

	_mysql_executor_container[key] = mysql_executor
	return mysql_executor
}

func GetMysqlExecutor(key string) *T_Mysql_Executor {
	return _mysql_executor_container[key]
}

/**
* @brief select xxx from xxx_table where xxxx
*
* @param T_Mysql_Executor
 */
func (mysql_executor *T_Mysql_Executor) Query(sql string, values ...interface{}) (I_Mysql_Rows, error) {
	stmt, err := mysql_executor.db.Prepare(sql)
	if nil != err {
		return nil, err
	}

	mysql_executor.logSql(sql, values...)
	rows, errq := stmt.Query(values...)
	if nil != errq {
		return nil, errq
	}

	return rows, nil
}

/**
* @brief select xxx from xxx_table where xxx, only one row
*
* @param T_Mysql_Executor
 */
func (mysql_executor *T_Mysql_Executor) QueryRow(sql string, values ...interface{}) (I_Mysql_Row, error) {
	stmt, err := mysql_executor.db.Prepare(sql)
	if nil != err {
		return nil, err
	}

	mysql_executor.logSql(sql, values...)
	return stmt.QueryRow(values...), nil
}

/**
* @brief insert xxx, update xxx
*
* @param T_Mysql_Executor
 */
func (mysql_executor *T_Mysql_Executor) Exec(sql string, values ...interface{}) (I_Mysql_Result, error) {
	stmt, err := mysql_executor.db.Prepare(sql)
	if nil != err {
		return nil, err
	}

	mysql_executor.logSql(sql, values...)
	result, erre := stmt.Exec(values...)
	if nil != erre {
		return nil, erre
	}

	return result, nil
}

func (mysql_executor *T_Mysql_Executor) logSql(sql string, values ...interface{}) {
	if nil == mysql_executor.logger {
		return
	}

	sql = strings.Replace(sql, "?", "%s", -1)
	vs := make([]interface{}, len(values))

	for i, v := range values {
		s := fmt.Sprint(v)
		switch v.(type) {
		case int:
			vs[i] = s
		case string:
			vs[i] = "'" + s + "'"
		}
	}

	sql_str := fmt.Sprintf(sql, vs...)
	mysql_executor.logger.Log(sql_str)
}
