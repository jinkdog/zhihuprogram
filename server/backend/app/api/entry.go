package api

import (
	"main/app/api/answer"
	"main/app/api/article"
	"main/app/api/comment"
	"main/app/api/question_community"
	"main/app/api/user"
)

var insUser = user.Group{}
var insQuestion = question_community.Group{}
var insAnswer = answer.Group{}
var insComment = comment.Group{}
var insArticle = article.Group{}

func User() *user.Group {
	return &insUser
}

func Question() *question_community.Group {
	return &insQuestion
}
func Answer() *answer.Group {
	return &insAnswer
}
func Comment() *comment.Group {
	return &insComment
}
func Article() *article.Group {
	return &insArticle
}
