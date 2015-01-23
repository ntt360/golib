/**
* @file conf.go
* @brief log_conf
* @author ligang
* @version 1.0
* @date 2014-10-17
 */

package bizlog

import (
	"github.com/mydoraemon/golib/biztime"
)

type t_log_conf struct {
	r_path  string
	split   int
	suffix  string
	bufsize int
}

/**
* @brief new log_conf
*
* @param string
* @param string
* @param int
* @param int
*
* @return *t_log_conf
 */
func newLogConf(r_path string, split int, bufsize int) *t_log_conf {
	return &t_log_conf{
		r_path:  r_path,
		split:   split,
		suffix:  makeFileSuffix(split),
		bufsize: bufsize,
	}
}

/**
* @brief suffix by split
*
* @param int
*
* @return suffix
 */
func makeFileSuffix(split int) string {
	switch split {
	case SPLIT_BY_DAY:
		return biztime.Now().Format(biztime.FMT_STR_YEAR + biztime.FMT_STR_MONTH + biztime.FMT_STR_DAY)
	case SPLIT_BY_HOUR:
		return biztime.Now().Format(biztime.FMT_STR_YEAR + biztime.FMT_STR_MONTH + biztime.FMT_STR_DAY + biztime.FMT_STR_HOUR)
	default:
		return ""
	}
}
