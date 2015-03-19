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
)

func (executor *T_Redis_Executor) Expire(key string, expire int) error {
	_, err := executor.conn.Do("EXPIRE", key, expire)

	return err
}

func (executor *T_Redis_Executor) Pexpire(key string, pexpire int) error {
	_, err := executor.conn.Do("PEXPIRE", key, pexpire)

	return err
}

func (executor *T_Redis_Executor) Del(keys ...string) error {
	args := make([]interface{}, 0, len(keys))
	for _, k := range keys {
		args = append(args, k)
	}
	_, err := executor.conn.Do("DEL", args...)

	return err
}

func (executor *T_Redis_Executor) Rename(old_key string, new_key string) error {
	_, err := executor.conn.Do("RENAME", old_key, new_key)

	return err
}
