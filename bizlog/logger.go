/**
* @file logger.go
* @brief logger
* @author ligang
* @version 1.0
* @date 2014-12-25
 */

package bizlog

import (
	"sync"
)

var _logger I_Logger
var _logger_container map[string]I_Logger

type I_Logger interface {
	Log(msg string)
	Flush()
	Free()

	writeMsg(msg string)
}

/**
* @name base logger
* @{ */

type t_base_logger struct {
	key      string
	conf     *t_log_conf
	formater i_msg_formater
	writer   i_log_writer
	lock     *sync.Mutex
}

func (logger *t_base_logger) Flush() {
	logger.lock.Lock()
	logger.writer.flush()
	logger.lock.Unlock()
}

func (logger *t_base_logger) Free() {
	logger.lock.Lock()
	logger.writer.free()
	logger.lock.Unlock()
}

func (logger *t_base_logger) writeMsg(msg string) {
	msg = logger.formater.fmt(msg)
	logger.checkWriter()

	logger.lock.Lock()
	logger.writer.writeln(msg)
	logger.lock.Unlock()
}

func (logger *t_base_logger) checkWriter() {
	suffix := makeFileSuffix(logger.conf.split)

	if suffix != logger.conf.suffix {
		path := makeLogPath(logger.key, logger.conf.r_path, suffix)

		logger.lock.Lock()
		logger.conf.suffix = suffix
		logger.writer.free()
		logger.writer = newLogWriter(path, logger.conf.bufsize)
		logger.lock.Unlock()
	}
}

func newBaseLogger(key string, conf *t_log_conf, formater i_msg_formater, writer i_log_writer) *t_base_logger {
	return &t_base_logger{
		key:      key,
		conf:     conf,
		formater: formater,
		writer:   writer,
		lock:     new(sync.Mutex),
	}
}

/**  @} */

/**
* @name sync logger
* @{ */

type t_sync_logger struct {
	*t_base_logger
}

func (logger *t_sync_logger) Log(msg string) {
	logger.writeMsg(msg)
}

/**  @} */

/**
* @name async logger
* @{ */

type t_async_logger struct {
	*t_base_logger
}

func (logger *t_async_logger) Log(msg string) {
	sendToLogRoutine(logger.key, msg)
}

// func (logger *t_async_logger) Flush() {
//     endLogRoutine()
//     logger.t_base_logger.Flush()
// }

/**  @} */
