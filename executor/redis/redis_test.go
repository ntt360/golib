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
	"testing"
)

/**
* @name string
* @{ */

func TestSet(t *testing.T) {
	executor := getExecutor()
	executor.Set("abc", "hello")
	executor.Pexpire("abc", 10*1000)
}

func TestGet(t *testing.T) {
	value, err, exist := getExecutor().Get("abc")
	fmt.Println(value, err, exist)
}

/**  @} */

/**
* @name set
* @{ */

func TestSadd(t *testing.T) {
	executor := getExecutor()
	executor.Sadd("tset1", "vbox01", "vbox02")
	executor.Sadd("tset2", "vbox03", "vbox04")

	executor.Expire("tset1", 10)
	executor.Pexpire("tset2", 10*1000)
}

func TestSmembers(t *testing.T) {
	executor := getExecutor()
	values, err, exist := executor.Smembers("tset1")
	fmt.Println(values, err, exist)
}

func TestSunion(t *testing.T) {
	executor := getExecutor()
	values, err, exist := executor.Sunion("tset1", "tset2")
	fmt.Println(values, err, exist)
}

func TestDel(t *testing.T) {
	executor := getExecutor()
	err := executor.Del("abc")
	fmt.Println(err)
}

func TestRename(t *testing.T) {
	executor := getExecutor()
	err := executor.Rename("tset2", "tset1")
	fmt.Println(err)
}

/**  @} */

func getExecutor() *T_Redis_Executor {
	redis_conf := T_Redis_Conf{
		Host: "127.0.0.1",
		Port: "6379",
		Pass: "123",
	}

	executor, _ := NewExecutor(redis_conf)
	return executor
}
