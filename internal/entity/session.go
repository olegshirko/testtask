package entity

import "time"

type Session struct {
	Id      string    `db:"id"`
	Uid     int64     `db:"uid"`
	Created time.Time `db:"created_at"`
	UserIp  string    `db:"user_ip"`
}
