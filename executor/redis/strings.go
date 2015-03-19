/**
* @file strings.go
* @brief strings
* @author ligang
* @version 1.0
* @date 2015-03-19
 */

package redis

import (
	//     "fmt"
	redigo "github.com/garyburd/redigo/redis"
)

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
