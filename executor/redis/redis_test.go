/**
* @file redis_test.go
* @brief test redis executor
* @author ligang
* @version 1.0
* @date 2015-03-18
 */

package redis

import (
	"fmt"
	"strconv"
	"testing"
)

func TestSet(t *testing.T) {
	err := getExecutor().Set("abc", strconv.Itoa(1234))
	err = getExecutor().Set("abc", "hello")
	fmt.Println(err)
}

func TestGet(t *testing.T) {
	value, err, exist := getExecutor().Get("abc")
	fmt.Println(value, err, exist)
}

func getExecutor() *T_Redis_Executor {
	redis_conf := T_Redis_Conf{
		Host: "127.0.0.1",
		Port: "6379",
		Pass: "123",
	}

	executor, _ := NewExecutor(redis_conf)
	return executor
}
