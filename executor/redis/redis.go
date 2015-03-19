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
	redigo "github.com/garyburd/redigo/redis"
	"time"
)

type T_Redis_Conf struct {
	Host string
	Port string
	Pass string
}

type T_Redis_Executor struct {
	conn redigo.Conn
}

func NewExecutor(redis_conf T_Redis_Conf) (*T_Redis_Executor, error) {
	executor := new(T_Redis_Executor)

	address := redis_conf.Host + ":" + redis_conf.Port
	connect_timeout, _ := time.ParseDuration(DEF_CONNECT_TIMEOUT)
	read_timeout, _ := time.ParseDuration(DEF_READ_TIMEOUT)
	write_timeout, _ := time.ParseDuration(DEF_WRITE_TIMEOUT)

	conn, err := redigo.DialTimeout("tcp", address, connect_timeout, read_timeout, write_timeout)
	if nil != err {
		return nil, err
	}

	executor.conn = conn
	_, err = executor.conn.Do("AUTH", redis_conf.Pass)
	if nil != err {
		return nil, err
	}

	return executor, nil
}
