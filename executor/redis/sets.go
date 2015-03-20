/**
* @file sets.go
* @brief sets
* @author ligang
* @version 1.0
* @date 2015-03-19
 */

package redis

import (
	//     "fmt"
	redigo "github.com/garyburd/redigo/redis"
)

func (executor *T_Redis_Executor) Sadd(key string, values ...string) error {
	args_str := key
	args := make([]interface{}, 0, len(values)+1)
	args = append(args, key)
	for _, v := range values {
		args_str += " " + v
		args = append(args, v)
	}

	cmd := "SADD"
	executor.logDo(cmd, args_str)
	_, err := executor.conn.Do(cmd, args...)

	return err
}

func (executor *T_Redis_Executor) Smembers(key string) ([]string, error, bool) {
	cmd := "SMEMBERS"
	executor.logDo(cmd, key)
	reply, err := executor.conn.Do(cmd, key)

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

	cmd := "SUNION"
	executor.logDo(cmd, keys...)
	reply, err := executor.conn.Do(cmd, args...)

	if nil != err || nil == reply {
		return []string{}, err, false
	}

	values, err := redigo.Strings(reply, err)
	return values, err, true
}
