package util

import "time"

func Str2time(t string) time.Time {
	tz, _ := time.LoadLocation("Asia/Tokyo")
	timeJST, _ := time.ParseInLocation("2006-01-02", t, tz)
	return timeJST
}
