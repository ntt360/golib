/**
* @file const.go
* @brief const
* @author ligang
* @version 1.0
* @date 2014-12-25
 */

package bizlog

const (
	MODE_SYNC  = 1
	MODE_ASYNC = 2
)

const (
	SPLIT_NO      = 0
	SPLIT_BY_DAY  = 1
	SPLIT_BY_HOUR = 2
)

const (
	MSG_FMT_PURE        = 1
	MSG_FMT_LINE_HEADER = 2
)

const (
	DEF_COL_SPR = "\t"
	DEF_BUFSIZE = 4096
)

const CAP_LOG_QUEUE = 1024

/**
* @name buildin log key
* @{ */

const (
	LOG_KEY_ERROR = "error"
)

/**  @} */
