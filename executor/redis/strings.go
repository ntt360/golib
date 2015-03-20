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
	cmd := "SET"
	executor.logDo(cmd, key, value)
	_, err := executor.conn.Do(cmd, key, value)

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
	cmd := "GET"
	executor.logDo(cmd, key)
	reply, err := executor.conn.Do(cmd, key)

	if nil != err || nil == reply {
		return "", err, false
	}

	value, err := redigo.String(reply, err)
	return value, err, true
}
