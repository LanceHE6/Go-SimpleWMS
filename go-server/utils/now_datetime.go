package utils

import "time"

func GetNowDateTime() string {
	// 获取当前时间
	now := time.Now()
	// 格式化为"2006-01-02 15:04:05"的字符串
	nowDateTime := now.Format("2006-01-02 15:04:05")
	return nowDateTime
}
