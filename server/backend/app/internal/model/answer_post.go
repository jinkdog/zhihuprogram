package model

import "time"

type AnswerPost struct {
	Id                  int64     `json:"id" form:"id" db:"id"`
	AnswerId            int64     `json:"answer_id" form:"answer_id" db:"answer_id"`
	Title               string    `json:"title" form:"title" db:"title"`
	Content             string    `json:"content" form:"content" db:"content"`
	AuthorId            int64     `json:"author_id" form:"author_id" db:"author_id"`
	QuestionCommunityId int64     `json:"question_community_id" form:"question_community_id" db:"question_community_id"`
	Status              int64     `json:"status" form:"status" db:"status"`
	CreateTime          time.Time `json:"create_time" form:"create_time" db:"create_time"`
	UpdateTime          time.Time `json:"update_time" form:"update_time" db:"update_time"`
}
