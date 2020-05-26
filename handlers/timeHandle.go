package handlers

import (
	"log"
	"remind-go/models"
	"time"
)

var todos [] models.Todo

var isFirst bool = true

func Scheduler() {
	//3小时醒过来一次查看下一个三小时内要发送的短信
	for {
		if isFirst == false {
			timer := time.NewTicker(3 * time.Hour)
			<-timer.C
		}
		isFirst = false
		now := GetLocalTimeNow()
		hh, _ := time.ParseDuration("1h")
		threeTime := now.Add(hh * 3)
		//查询最近三小时内有没有要发送的短信
		rows, err := models.Db.Query(
			"select * from todos where notice_time>? and notice_time<? and status=?",
			SetFormatTime(now), SetFormatTime(threeTime), 2)
		if err != nil {
			log.Println(err.Error())
		}
		for rows.Next() {
			var todo = models.Todo{}
			var email = &Email{}
			var phone = &Phone{}
			if err = rows.Scan(&todo.Id, &todo.Content, &todo.CreatedAt,
				&todo.NoticeTime, &todo.Status, &todo.Phone, &todo.Email); err != nil {
				log.Println(err.Error())
			}
			email.Body = todo.Content
			phone.Phone = todo.Phone
			phone.Id = todo.Id
			go func(todo2 models.Todo, email2 *Email, phone2 *Phone) {
				SendEmailOrPhone(todo2, email2, phone2)
			}(todo, email, phone)
			todos = append(todos, todo)
		}

	}
}
