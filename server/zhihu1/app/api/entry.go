package api

import (
	"main/app/api/question_community"
	"main/app/api/user"
)

var insUser = user.Group{}
var insQuestion = question_community.Group{}

func User() *user.Group {
	return &insUser
}

func Question() *question_community.Group {
	return &insQuestion
}
