package models

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

type Todo struct {
	Id         int64
	Content    string
	CreatedAt  time.Time
	NoticeTime time.Time
	Status     int
	Phone      string
	Email      sql.NullString
}

func CreateToDo(createTime time.Time, content string, phone string, sendTime string) (lastId int64, err error) {
	statemt := "insert into todos(content,created_at,notice_time,status,phone,email)values(?,?,?,?,?,?)"
	stmtin, err := Db.Prepare(statemt)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer stmtin.Close()
	res, err := stmtin.Exec(content, createTime.Format("2006-01-02 15:04:05"), sendTime, 2, phone, "")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	lastInsertId, err := res.LastInsertId()
	return lastInsertId, err
}

//短信发送成功 标识一下 不要重复发送
func SetSuccessStatus(id int64) {
	statement := "update todos set status=? where id=?"
	stim, err := Db.Prepare(statement)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer stim.Close()
	_, err = stim.Exec(1, id)
	if err != nil {
		log.Println(err.Error())
	}
	return
}
