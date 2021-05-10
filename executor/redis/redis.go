/**
* @file redis.go
* @brief redis executor
* @author ligang
* @version 1.0
* @date 2015-03-18
 */

package redis

import (
	//     "fmt"
	"strings"
	"time"

	redigo "github.com/garyburd/redigo/redis"
	"github.com/ntt360/golib/bizlog"
)

type T_Redis_Conf struct {
	Host string
	Port string
	Pass string
}

type T_Redis_Executor struct {
	conn   redigo.Conn
	logger bizlog.I_Logger
}

func NewExecutor(redis_conf T_Redis_Conf, logger bizlog.I_Logger) (*T_Redis_Executor, error) {
	executor := new(T_Redis_Executor)

	address := redis_conf.Host + ":" + redis_conf.Port
	connect_timeout, _ := time.ParseDuration(DEF_CONNECT_TIMEOUT)
	read_timeout, _ := time.ParseDuration(DEF_READ_TIMEOUT)
	write_timeout, _ := time.ParseDuration(DEF_WRITE_TIMEOUT)

	conn, err := redigo.DialTimeout("tcp", address, connect_timeout, read_timeout, write_timeout)
	if nil != err {
		return nil, err
	}

	if "" != redis_conf.Pass {
		_, err = conn.Do("AUTH", redis_conf.Pass)
		if nil != err {
			return nil, err
		}
	}

	executor.conn = conn
	executor.logger = logger

	return executor, nil
}

func (executor *T_Redis_Executor) Close() error {
	return executor.conn.Close()
}

func (executor *T_Redis_Executor) logDo(cmd string, args ...string) {
	if nil == executor.logger {
		return
	}

	msg := cmd + " " + strings.Join(args, " ")
	executor.logger.Log(msg)
}
