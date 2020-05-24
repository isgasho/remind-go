package handlers

import "time"

func GetLocalTimeNow() time.Time {
	sh, _ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(sh)
}

func SetFormatTime(timeDate time.Time) string {
	return timeDate.Format("2006-01-02 15:04:05")
}

