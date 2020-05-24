package handlers

import (
	"fmt"
	"log"
	"regexp"
	"remind-go/models"
	"time"
)

var timeDay = map[string]string{
	"今天":  getDateString(0),
	"明天":  getDateString(1),
	"后天":  getDateString(2),
	"大后天": getDateString(3),
}

//var timeHMS = map[string]string{
//	"个月": getDateString(0),
//	"小时": getDateString(1),
//	"分钟": getDateString(2),
//	"分":  getDateString(2),
//	"秒":  getDateString(3),
//	"周":  getDateString(3),
//	"天":  getDateString(3),
//}

type contentRegexp struct {
	*regexp.Regexp
}

//计算日期
func getDateString(count int) string {
	t := time.Now()
	newTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	//通知时间
	noticeTime := newTime.AddDate(0, 0, count)
	logDay := noticeTime.Format("2006-01-02")
	return logDay
}

//时间匹配
var myexp = contentRegexp{regexp.MustCompile(
	`(今天|明天|后天|大后天|[\d]{4}-[\d]{2}-[\d]{2}\s[\d]{2}:[\d]{2}|[\d]{8}\s[\d]{1,2}:[\d]{1,2}|[[\d]{1,2}:[\d]{1,2}|[\d]{1,2}(个月|小时|点|分钟|分|秒|周|天))`,
)}

//手机号匹配
var phone = contentRegexp{regexp.MustCompile(
	`(1[356789]\d)(\d{4})(\d{4})`,
)}

func HandleMessage(content string) string {
	phone := phone.FindStringSubmatch(content)
	if phone == nil {
		return "不留下联系方式我咋么联系上您"
	}
	//手机号
	fmt.Println(phone[0])
	mmp := myexp.FindAllStringSubmatch(content, -1)
	fmt.Println(mmp)
	if mmp == nil {
		return "小姐姐，你这个时间格式有点为难我了"
	}
	//最多只有三位 时 分 秒
	if len(mmp) > 3 {
		mmp = mmp[:3]
	}

	var realDate string
	for _, item := range mmp {
		//今天明天后台大后天
		if _, ok := timeDay[item[0]]; ok {
			realDate = timeDay[item[0]]
		} else {
			if realDate == "" {
				realDate = item[0]
			} else {
				realDate = realDate + " " + item[0]
			}
		}
	}
	fmt.Println(realDate)

	cstSh, _ := time.LoadLocation("Asia/Shanghai")
	createdTime := time.Now().In(cstSh)

	var lastId int64
	lastId, err := models.CreateToDo(createdTime, content, phone[0], realDate)
	if err != nil {
		log.Println(err.Error())
	}
	isCreateTimerForSendNotice(lastId, realDate, createdTime, phone[0])
	return "ahha"
}

//通知时间小于现在的3小时，直接搞个定时器
func isCreateTimerForSendNotice(lastId int64, sendTime string, createdTime time.Time, phone string) {
	log.Println(sendTime)
	cstSh, _ := time.LoadLocation("Asia/Shanghai")
	noticeTime, err := time.ParseInLocation("2006-01-02 15:04:05", sendTime+":00", cstSh)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	diff := noticeTime.Sub(createdTime)
	//直接给他一个定时器 执行 即使下一个断续器启动 检索信息的时候这条通知已经标注已通知了
	if diff.Hours() < 3 && diff.Hours() > 0 {
		var noticePhone = &Phone{}
		noticePhone.Phone = phone
		go func() {
			//到点执行
			timer := time.NewTimer(diff)
			<-timer.C
			noticePhone.SendNotice(lastId)
		}()
	}
}
