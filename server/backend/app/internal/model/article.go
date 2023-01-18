package model

import "time"

type Article struct {
	Id         int64     `json:"id" form:"id" db:"id" `
	ArticleId  int64     `json:"article_id" form:"article_id" db:"article_id" `
	Title      string    `json:"title" form:"title" db:"title" `
	Content    string    `json:"content" form:"content" db:"content" `
	AuthorId   int64     `json:"author_id" form:"author_id" db:"author_id" `
	CreateTime time.Time `json:"create_time" form:"create_time" db:"create_time" `
	UpdateTime time.Time `json:"update_time" form:"update_time" db:"update_time" `
}
