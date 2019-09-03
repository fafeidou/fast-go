package util

import (
	"time"
)

//时间转时间戳
func DateTimeToTimestamp(datetime string) int64 {
	loc, _ := time.LoadLocation("Asia/Shanghai") //设置时区
	tt, err := time.ParseInLocation("2006-01-02 15:04:05", datetime, loc)
	if err != nil {
		return 0
	}
	return tt.Unix()
}

//时间转时间戳
func DateToTimestamp(datetime string) int64 {
	loc, _ := time.LoadLocation("Asia/Shanghai") //设置时区
	tt, err := time.ParseInLocation("2006-01-02", datetime, loc)
	if err != nil {
		return 0
	}
	return tt.Unix()
}

//时间戳 to 时间
func TimestampToDate(timestamp int64) string {
	return time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
}

//当前时间
func NowDatetime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// 当前时间戳
func NowTimestamp() int64 {
	return time.Now().Unix()
}
