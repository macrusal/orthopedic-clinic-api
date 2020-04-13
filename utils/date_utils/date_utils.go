package date_utils

import (
	"time"
)

const(
	apiDateLayout = "2006-02-01T15:04:05"
)

func GetNow() time.Time {
	return time.Now()
}

func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}