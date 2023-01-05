package service

import (
	"main/app/internal/service/question_community"
	"main/app/internal/service/user"
)

var insUser = user.Group{}
var insQuestion = question_community.Group{}

func User() *user.Group {
	return &insUser
}

func Question() *question_community.Group {
	return &insQuestion
}
