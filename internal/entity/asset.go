package entity

import "time"

type Asset struct {
	Name    string    `json:"name"`
	Uid     int64     `json:"uid"`
	Data    []byte    `json:"data"`
	Created time.Time `json:"created"`
}
