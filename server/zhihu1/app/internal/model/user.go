package model

import "time"

type UserSubject struct {
	Id         int64     `json:"id" form:"id" db:"id"`
	Username   string    `json:"username" form:"username" db:"username"`
	Password   string    `json:"password" form:"password" db:"password"`
	CreateTime time.Time `json:"create_time" form:"create_time" db:"create_time"`
	UpdateTime time.Time `json:"update_time" form:"update_time" db:"update_time"`
}
