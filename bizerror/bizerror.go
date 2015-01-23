/**
* @file bizerror.go
* @brief error struct
* @author ligang
* @version 1.0
* @date 2014-12-30
 */

package bizerror

import (
	"strconv"
)

type T_Error struct {
	errno int
	msg   string
}

func NewError(errno int, msg string) *T_Error {
	return &T_Error{
		errno: errno,
		msg:   msg,
	}
}

func (err *T_Error) Error() string {
	result := "errno: " + strconv.Itoa(err.errno) + ", "
	result += "msg: " + err.msg

	return result
}

func (err *T_Error) GetErrno() int {
	return err.errno
}

func (err *T_Error) GetMsg() string {
	return err.msg
}
