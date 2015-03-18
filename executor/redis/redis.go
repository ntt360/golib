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

/**
* @name string
* @{ */

func (executor *T_Redis_Executor) Set(key string, value string) error {
	_, err := executor.conn.Do("SET", key, value)

	return err
}

/**
* @brief get string
*
* @param key string
*
* @return value, error, exist
 */
func (executor *T_Redis_Executor) Get(key string) (string, error, bool) {
	reply, err := executor.conn.Do("GET", key)

	if nil != err || nil == reply {
		return "", err, false
	}

	value, err := redigo.String(reply, err)
	return value, err, true
}

/**  @} */

/**
* @name set
* @{ */

func (executor *T_Redis_Executor) Sadd(key string, values ...string) error {
	args := make([]interface{}, 0, len(values)+1)
	args = append(args, key)
	for _, v := range values {
		args = append(args, v)
	}
	_, err := executor.conn.Do("SADD", args...)

	return err
}

func (executor *T_Redis_Executor) Smembers(key string) ([]string, error, bool) {
	reply, err := executor.conn.Do("SMEMBERS", key)

	if nil != err || nil == reply {
		return []string{}, err, false
	}

	values, err := redigo.Strings(reply, err)
	return values, err, true
}

func (executor *T_Redis_Executor) Sunion(keys ...string) ([]string, error, bool) {
	args := make([]interface{}, 0, len(keys))
	for _, k := range keys {
		args = append(args, k)
	}
	reply, err := executor.conn.Do("SUNION", args...)

	if nil != err || nil == reply {
		return []string{}, err, false
	}

	values, err := redigo.Strings(reply, err)
	return values, err, true
}

/**  @} */
