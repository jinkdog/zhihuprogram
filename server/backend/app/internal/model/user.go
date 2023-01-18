package model

import "time"

type UserSubject struct {
	Id         int64     `json:"id" form:"id" db:"id"`
	Username   string    `json:"username" form:"username" db:"username"`
	Password   string    `json:"password" form:"password" db:"password"`
	CreateTime time.Time `json:"create_time" form:"create_time" db:"create_time"`
	UpdateTime time.Time `json:"update_time" form:"update_time" db:"update_time"`
}

//存放UserSubject结构体

type UserCollection struct {
	Id          int64     `json:"id" form:"id" db:"id"`
	UserId      int64     `json:"user_id" form:"user_id" db:"user_id"`
	CollectType int32     `json:"collect_type" form:"collect_type" db:"collect_type"`
	QuestionId  string    `json:"question_community_id" form:"question_community_id" db:"question_community_id"`
	AnswerId    int64     `json:"answer_post_id" form:"answer_post_id" db:"answer_post_id"`
	CreateTime  time.Time `gorm:"autoCreateTime" json:"create_time" form:"create_time" db:"create_time"`
	UpdateTime  time.Time `gorm:"autoUpdateTime" json:"update_time" form:"update_time" db:"update_time"`
}

type Collection struct {
	Id             int64       `json:"id"`
	CollectionType string      `json:"collection_type"`
	CollectionData interface{} `json:"collection_data"`
}
