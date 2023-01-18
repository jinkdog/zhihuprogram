package service

import (
	"main/app/internal/service/answer_post"
	"main/app/internal/service/article"
	"main/app/internal/service/comment"
	"main/app/internal/service/question_community"
	"main/app/internal/service/user"
)

var insUser = user.Group{}
var insQuestion = question_community.Group{}
var insAnswer = answer_post.Group{}
var insComment = comment.Group{}
var insArticle = article.Group{}

func User() *user.Group {
	return &insUser
}

func Question() *question_community.Group {
	return &insQuestion
}

func Answer() *answer_post.Group {
	return &insAnswer
}

func Comment() *comment.Group {
	return &insComment
}
func Article() *article.Group {
	return &insArticle
}
