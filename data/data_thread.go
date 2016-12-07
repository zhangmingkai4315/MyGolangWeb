package data

import (
	"time"
)

type Thread struct {
	Id           int
	Uuid         string
	Title        string
	Content      string
	TypeOfThread string
	Status       string
	UserId       string
	CreatedAt    time.Time
	ModifiedAt   time.Time
}

func (thread *Thread) NumReplies() (count int) {
	rows, err := Db.Query("select count(*) from comments where thread_id =$1", thread.Id)
	if err != nil {
		return
	}
	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			return
		}
	}
	rows.Close()
	return
}

func Threads() (threads []Thread, err error) {
	rows, err := Db.Query("select id,uuid,title,user_id,created_at,modification_time from threads order by created_at desc")
	if err != nil {
		return
	}
	for rows.Next() {
		th := Thread{}
		if err = rows.Scan(&th.Id, &th.Uuid, &th.Title, &th.UserId, &th.CreatedAt, &th.ModifiedAt); err != nil {
			return
		}
		threads = append(threads, th)
	}
	rows.Close()
	return
}
