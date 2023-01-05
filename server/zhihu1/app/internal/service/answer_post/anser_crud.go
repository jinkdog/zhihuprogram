package answer_post

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	g "main/app/global"
	"main/app/internal/model"
)

type SAnswer struct{}

var insAnswer = SAnswer{}

func (s *SAnswer) GetAnswerList(ctx context.Context) (answer []*model.AnswerPost, err error) {
	rows, err := g.MysqlDB.QueryContext(ctx, "select answer_id,title,content,author_id,question_community_id,create_time,update_time from anwer_post  order by create_time")

	for rows.Next() {
		if err != nil {
			if err != rows.Scan(&answer) {
				g.Logger.Error("query mysql record failed.",
					zap.Error(err),
					zap.String("table", "answer_post"),
				)
				return nil, fmt.Errorf("internal err")
			} else {
				return nil, fmt.Errorf("no answer_post in the db")
			}
		}
	}
	return answer, nil
}

func (s *SAnswer) GetAnswerListById(id uint64, ctx context.Context) error {
	answer := new(model.AnswerPost)
	rows, err := g.MysqlDB.QueryContext(ctx, "select answer_id,title,content,author_id,question_community_id from answer_post where answer_id=?", id)

	for rows.Next() {
		if err != nil {
			if err != rows.Scan(&answer.AnswerId, &answer.Title, &answer.Content, &answer.AuthorId, &answer.QuestionCommunityId) {
				g.Logger.Error("query mysql record failed.",
					zap.Error(err),
					zap.String("table", "answer_post"),
				)
				return fmt.Errorf("internal err")
			} else {
				return fmt.Errorf("no answer_post in the db")
			}
		}
	}
	return nil
}

//func (s *SAnswer) CreateAnswer(ctx context.Context) error {
//	answer := new(model.AnswerPost)
//	_, err := g.MysqlDB.ExecContext(ctx)
//}
