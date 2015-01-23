/**
* @file biztime.go
* @brief time knowledge
* @author ligang
* @version 1.0
* @date 2014-10-15
 */

package biztime

import (
	"time"
)

const (
	FMT_STR_YEAR   = "2006"
	FMT_STR_MONTH  = "01"
	FMT_STR_DAY    = "02"
	FMT_STR_HOUR   = "15"
	FMT_STR_MINUTE = "04"
	FMT_STR_SECOND = "05"
)

type I_Time interface {
	IsZero() bool
	Unix() int64
	Format(layout string) string
}

func Now() I_Time {
	return time.Now()
}

func Parse(layout string, value string) I_Time {
	ts, _ := time.Parse(layout, value)

	return ts
}

func GetGeneralLayout() string {
	layout := FMT_STR_YEAR + "-" + FMT_STR_MONTH + "-" + FMT_STR_DAY + " "
	layout += FMT_STR_HOUR + ":" + FMT_STR_MINUTE + ":" + FMT_STR_SECOND

	return layout
}
