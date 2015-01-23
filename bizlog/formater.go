/**
* @file msg.go
* @brief log msg
* @author ligang
* @version 1.0
* @date 2014-12-25
 */

package bizlog

import (
	"mydoraemon/golib/biztime"
)

type i_msg_formater interface {
	fmt(msg string) string
}

/**
* @name add line_header
* @{ */

type t_line_header_formater struct {
}

func (formater *t_line_header_formater) fmt(msg string) string {
	result := "[" + biztime.Now().Format(biztime.FMT_STR_YEAR+"-"+biztime.FMT_STR_MONTH+"-"+biztime.FMT_STR_DAY+" "+biztime.FMT_STR_HOUR+":"+biztime.FMT_STR_MINUTE+":"+biztime.FMT_STR_SECOND) + "]"
	result += DEF_COL_SPR + msg

	return result
}

/**  @} */

/**
* @name pure msg
* @{ */

type t_pure_formater struct {
}

func (formater *t_pure_formater) fmt(msg string) string {
	return msg
}

/**  @} */

func newMsgFormator(key int) i_msg_formater {
	if MSG_FMT_LINE_HEADER == key {
		return new(t_line_header_formater)
	} else {
		return new(t_pure_formater)
	}
}
