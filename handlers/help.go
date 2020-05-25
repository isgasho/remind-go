package handlers

import (
	"strconv"
	"time"
)

func GetLocalTimeNow() time.Time {
	sh, _ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(sh)
}

func SetFormatTime(timeDate time.Time) string {
	return timeDate.Format("2006-01-02 15:04:05")
}

//浮点数转字符串截取
func Decimal(value float64) string {
	string := strconv.FormatFloat(value, 'f', 6, 64)
	return string[:4]
}
