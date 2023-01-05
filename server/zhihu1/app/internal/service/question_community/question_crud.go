package question_community

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	g "main/app/global"
	"main/app/internal/model"
	"time"
)

type SQuestion struct{}

var insQuestion = SQuestion{}

func (s *SQuestion) GetQuestionList(ctx context.Context) (question []*model.QuestionCommunity, err error) {
	//question = new([]model.QuestionCommunity)
	rows, err := g.MysqlDB.QueryContext(ctx, "select question_community_id,question_community_name from question_community")

	for rows.Next() {
		if err != nil {
			if err != rows.Scan(&question) {
				g.Logger.Error("query mysql record failed.",
					zap.Error(err),
					zap.String("table", "question_community"),
				)
				return nil, fmt.Errorf("internal err")
			} else {
				return nil, fmt.Errorf("no question_community in the db")
			}
		}
	}
	return question, nil
}

func (s *SQuestion) GetQuestionNameById(ctx context.Context, idStr string) (question *model.QuestionCommunity, err error) {
	question = new(model.QuestionCommunity)
	rows, err := g.MysqlDB.QueryContext(ctx, "select question_community_id,question_community_name from question_community where question_community_id=?", idStr)
	for rows.Next() {
		if err != nil {
			if err != rows.Scan(&question.QuestionCommunityId, &question.QuestionCommunityName) {
				g.Logger.Error("query mysql record failed",
					zap.Error(err),
					zap.String("table", "question_community"),
				)
				return nil, fmt.Errorf("internal err")
			}
		} else {
			return nil, fmt.Errorf("invalid question_community_id")
		}
	}
	return question, nil
}

func (s *SQuestion) GetQuestionByID(id uint64, ctx context.Context) (question *model.QuestionCommunity, err error) {
	question = new(model.QuestionCommunity)
	rows, err := g.MysqlDB.QueryContext(ctx, "select question_community_id,question_community_name,introduction,create_time from question_community where question_community_id=?", id)
	for rows.Next() {
		if err != nil {
			if err != rows.Scan(&question.QuestionCommunityId, &question.QuestionCommunityName, &question.Introduction, &question.CreateTime) {
				g.Logger.Error("query mysql record failed",
					zap.Error(err),
					zap.String("table", "question_community"),
				)
				return nil, fmt.Errorf("internal err")
			} else {
				return nil, fmt.Errorf("invalid question_community_id")
			}
		}
	}
	return question, nil
}

func (s *SQuestion) CreateQuestion(ctx context.Context) error {
	question := new(model.QuestionCommunity)
	_, err := g.MysqlDB.ExecContext(ctx, "insert into question_community(id,question_community_id,question_community_name,introduction,create_time,update_time)values (?,?,?,?,?,?)",
		question.Id,
		question.QuestionCommunityId,
		question.QuestionCommunityName,
		question.Introduction,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		g.Logger.Error("insert mysql record failed.",
			zap.Error(err),
			zap.String("table", "question_community"))
		return fmt.Errorf("insert mysql record failed")
	}
	return nil
}

func (s *SQuestion) DeleteQuestionByID(id uint64, ctx context.Context) error {
	//question := new(model.QuestionCommunity)
	_, err := g.MysqlDB.ExecContext(ctx, "delete from question_community where id=?", id)
	if err != nil {
		g.Logger.Error("delete mysql record failed.",
			zap.Error(err),
			zap.String("table", "question_community"))
		return fmt.Errorf("delete mysql record failed")
	}
	return nil
}

func (s *SQuestion) UpdateQuestionByID(id uint64, ctx context.Context) error {
	question := new(model.QuestionCommunity)
	_, err := g.MysqlDB.ExecContext(ctx, "update question_community set question_community_name=?,introduction=?,update_time=?where question_community_id=?",
		question.QuestionCommunityName,
		question.Introduction,
		time.Now(),
		id)
	if err != nil {
		g.Logger.Error("update mysql record failed.",
			zap.Error(err),
			zap.String("table", "question_community"))
		return fmt.Errorf("update mysal record failed")
	}
	return nil

}
