/**
* @file routine.go
* @brief log_routine
* @author ligang
* @version 1.0
* @date 2014-10-15
 */

package bizlog

import (
	"github.com/mydoraemon/golib/routine"
)

type t_async_log_msg struct {
	key string
	msg string
}

var _log_queue chan *t_async_log_msg
var _end_log routine.T_Routine_Status

/**
* @brief start routine
*
* @return
 */
func startLogRoutine() {
	_log_queue = make(chan *t_async_log_msg, CAP_LOG_QUEUE)
	_end_log = make(routine.T_Routine_Status)

	go func() {
		for {
			select {
			case async_log_msg, _ := <-_log_queue:
				logger := GetLogger(async_log_msg.key)
				if nil != logger {
					logger.writeMsg(async_log_msg.msg)
				}
			case <-_end_log:
				freeLogQueue()
				_end_log <- routine.ROUTINE_DONE
				return
			}
		}
	}()
}

/**
* @brief add log_msg to log_queue
*
* @param t_log_msg
*
* @return
 */
func sendToLogRoutine(key string, msg string) {
	async_log_msg := &t_async_log_msg{
		key: key,
		msg: msg,
	}
	_log_queue <- async_log_msg
}

/**
* @brief end routine
*
* @return
 */
func endLogRoutine() {
	_end_log <- routine.ROUTINE_START_END
	<-_end_log
}

/**
* @brief free log_queue
*
* @return
 */
func freeLogQueue() {
	for 0 != len(_log_queue) {
		async_log_msg, _ := <-_log_queue
		GetLogger(async_log_msg.key).writeMsg(async_log_msg.msg)
	}
}
