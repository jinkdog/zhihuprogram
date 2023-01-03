package model

import "time"

type QuestionCommunity struct {
	Id                    int64     `json:"id" form:"id" db:"id"`
	QuestionCommunityId   int64     `json:"question_community_id" form:"question_community_id" db:"question_community_id"`
	QuestionCommunityName string    `json:"question_community_name" form:"question_community_name" db:"question_community_name"`
	Introduction          string    `json:"introduction" form:"introduction" db:"introduction"`
	CreateTime            time.Time `json:"create_time" form:"create_time" db:"create_time"`
	UpdateTime            time.Time `json:"update_time" form:"update_time" db:"update_time"`
}
