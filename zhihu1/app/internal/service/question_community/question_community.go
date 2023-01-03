package question_community

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	g "main/app/global"
)

type SQuestion struct{}

var insQuestion = SQuestion{}

func (s *SQuestion) GetQuestionList(ctx context.Context) error {
	//question := &model.QuestionCommunity{}
	rows, err := g.MysqlDB.QueryContext(ctx, "select question_community,question_name from question_community")

	if err != nil {
		g.Logger.Error("query mysql record failed.",
			zap.Error(err),
			zap.String("table", "question_community"),
		)
		return fmt.Errorf("internal err")
	}

}
