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
