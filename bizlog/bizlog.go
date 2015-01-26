/**
* @file bizlog.go
* @brief bizlog
* @author ligang
* @version 1.0
* @date 2014-10-17
 */

package bizlog

import (
	"errors"
	//     "fmt"
	"github.com/mydoraemon/golib/tool"
)

var _log_root string

/**
* @brief must called first
*
* @param string
*
* @return
 */
func Init(log_root string) error {
	if !tool.DirExist(log_root) {
		return errors.New("log root not exists or not dir")
	}

	_log_root = log_root
	_logger_container = make(map[string]I_Logger)

	initBuildinLogger()
	startLogRoutine()

	return nil
}

/**
* @brief sync or async mode
*
* @param int
*
* @return
 */
func NewLogger(key string, mode int, r_path string, split int, bufsize int, msg_formater_key int) I_Logger {
	var logger I_Logger
	logger = GetLogger(key)
	if nil != logger {
		return logger
	}

	log_conf := newLogConf(r_path, split, bufsize)
	formater := newMsgFormator(msg_formater_key)
	path := makeLogPath(key, log_conf.r_path, log_conf.suffix)
	writer := newLogWriter(path, log_conf.bufsize)
	base_logger := newBaseLogger(key, log_conf, formater, writer)

	if MODE_ASYNC == mode {
		logger = &t_async_logger{
			base_logger,
		}
	} else {
		logger = &t_sync_logger{
			base_logger,
		}
	}

	_logger_container[key] = logger
	return logger
}

func GetLogger(key string) I_Logger {
	logger, ok := _logger_container[key]
	if ok {
		return logger
	} else {
		return nil
	}
}

/**
* @brief called for write log
*
* @param string
* @param string
*
* @return
 */
func Log(key string, msg string) {
	logger := GetLogger(key)
	if nil != logger {
		logger.Log(msg)
	}
}

/**
* @brief buildin error log
*
* @param string
*
* @return
 */
func Error(msg string) {
	Log(LOG_KEY_ERROR, msg)
}

/**
* @brief use when main done
*
* @return
 */
func Free() {
	endLogRoutine()
	for _, logger := range _logger_container {
		logger.Free()
	}
}

func initBuildinLogger() {
	NewLogger(LOG_KEY_ERROR, MODE_SYNC, "", SPLIT_NO, 0, MSG_FMT_LINE_HEADER)
}

/**
* @brief make log absolute path
*
* @param string
* @param string
* @param string
*
* @return
 */
func makeLogPath(key string, r_path string, suffix string) string {
	result := _log_root + "/"
	if "" != r_path {
		result += r_path + "/"
	}
	result += key
	result += ".log"
	if "" != suffix {
		result += "." + suffix
	}

	return result
}
