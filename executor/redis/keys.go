/**
* @file keys.go
* @brief keys
* @author ligang
* @version 1.0
* @date 2015-03-19
 */

package redis

import (
	//     "fmt"
	"strconv"
)

func (executor *T_Redis_Executor) Expire(key string, expire int) error {
	cmd := "EXPIRE"
	executor.logDo(cmd, key, strconv.Itoa(expire))
	_, err := executor.conn.Do(cmd, key, expire)

	return err
}

func (executor *T_Redis_Executor) Pexpire(key string, pexpire int) error {
	cmd := "PEXPIRE"
	executor.logDo(cmd, key, strconv.Itoa(pexpire))
	_, err := executor.conn.Do(cmd, key, pexpire)

	return err
}

func (executor *T_Redis_Executor) Del(keys ...string) error {
	args := make([]interface{}, 0, len(keys))
	for _, k := range keys {
		args = append(args, k)
	}

	cmd := "DEL"
	executor.logDo(cmd, keys...)
	_, err := executor.conn.Do(cmd, args...)

	return err
}

func (executor *T_Redis_Executor) Rename(old_key string, new_key string) error {
	cmd := "RENAME"
	executor.logDo(cmd, old_key, new_key)
	_, err := executor.conn.Do(cmd, old_key, new_key)

	return err
}
