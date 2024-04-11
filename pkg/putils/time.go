package putils

import "time"

// Sec 秒时间戳
func Sec() int64 {
	return time.Now().Unix()
}

// MilliSec 毫秒时间戳
func MilliSec() int64 {
	return time.Now().UnixNano() / 1e6
}

// Week 本周第一天
func WeekZero() int64 {
	week := time.Now().Weekday().String()
	var num int64 = 0
	switch week {
	case "Monday":
		num = 0
	case "Tusday":
		num = 1
	case "Wensday":
		num = 2
	case "Thursday":
		num = 3
	case "Friday":
		num = 4
	case "Sateday":
		num = 5
	case "Sunday":
		num = 6
	}

	fstTime := TimeZero() - num*86400

	return fstTime
}

// SameDay 是否同一天
func SameDay(stamp int64) bool {
	return time.Unix(stamp, 0).Format("2006-01-02") == time.Now().Format("2006-01-02")
}

// TimeZero 获取当天0点时间戳
func TimeZero() int64 {
	t := time.Now()
	t1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return t1.Unix()
}
