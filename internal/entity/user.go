package entity

import "time"

type User struct {
	Id        int64     `db:"id"`
	Username  string    `db:"login"`
	Password  string    `db:"password_hash"`
	CreatedAt time.Time `db:"created_at"`
}
