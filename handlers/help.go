package handlers

import (
	"log"
	"strconv"
	"sync"
	"time"
)

var ErrNoticeChannel = make(chan Phone, 10)

//错误的通知集合
var CountErr = make(map[int64]int, 10)

var lock sync.RWMutex

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

//处理发送失败通知
func HandlerErrNotice() {
	handles := getHandleErrNotice()
	for sendPhone := range handles {
		timer := time.NewTimer(time.Second * 30)
		<-timer.C
		if CountErr[sendPhone.Id] >= 3 {
			log.Printf("id为%d的通知重试超过三次了", sendPhone.Id)
			lock.RLock()
			delete(CountErr, sendPhone.Id)
			lock.RUnlock()
			continue
		}
		log.Println("重试中")
		lock.RLock()
		CountErr[sendPhone.Id] += 1
		lock.RUnlock()
		sendPhone.SendNotice(sendPhone.Id)
	}
}

//接收
func getHandleErrNotice() <-chan Phone {
	return ErrNoticeChannel
}

//发送
func SendHandleChannel(ch chan<- Phone, phone2 Phone) {
	ch <- phone2
}
