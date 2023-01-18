package model

import "time"

type Comment struct {
	Id         int64     `json:"id" form:"id" db:"id"`
	CommentId  int64     `json:"answer_id" form:"answer_id" db:"answer_id"`
	Content    string    `json:"content" form:"content" db:"content"`
	PostId     int64     `json:"post_id" form:"post_id" db:"post_id"`
	AuthorId   int64     `json:"author_id" form:"author_id" db:"author_id"`
	ParentId   int64     `json:"parent_id" form:"parent_id" db:"parent_id"`
	Status     int64     `json:"status" form:"status" db:"status"`
	CreateTime time.Time `json:"create_time" form:"create_time" db:"create_time"`
	UpdateTime time.Time `json:"update_time" form:"update_time" db:"update_time"`
}
