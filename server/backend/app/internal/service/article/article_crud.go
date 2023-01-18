package article

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	g "main/app/global"
	"main/app/internal/model"
)

type SArticle struct{}

var insArticle = SArticle{}

func (s *SArticle) GetArticleList(ctx context.Context) (article []*model.Article, err error) {
	rows, err := g.MysqlDB.QueryContext(ctx, "select article_id,title,content,author_id,create_time,update_time from article order by create_time")
	for rows.Next() {
		if err != nil {
			if err != rows.Scan(&article) {
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
	return article, nil
}

func (s *SArticle) GetArticleById(id uint64, ctx context.Context) error {
	article := new(model.Article)
	rows, err := g.MysqlDB.QueryContext(ctx, "select article_id,title,content,author_id from article where article_id=?", id)

	for rows.Next() {
		if err != nil {
			if err != rows.Scan(&article.ArticleId, &article.Title, &article.Content, &article.AuthorId) {
				g.Logger.Error("query mysql record failed.",
					zap.Error(err),
					zap.String("table", "article"),
				)
				return fmt.Errorf("internal err")
			} else {
				return fmt.Errorf("no article in the db")
			}
		}
	}
	return nil
}

func (s *SArticle) CreateArticle(ctx context.Context) error {
	article := new(model.AnswerPost)
	_, err := g.MysqlDB.ExecContext(ctx, "insert into article(article_id,title,content,author_id,create_time,update_time)values(?,?,?,?,?,?)",
		article.AnswerId,
		article.Title,
		article.Content,
		article.AuthorId,
		article.CreateTime,
		article.UpdateTime)

	if err != nil {
		g.Logger.Error("insert mysql record failed.",
			zap.Error(err),
			zap.String("table", "article"))
		return fmt.Errorf("insert mysql record failed")
	}
	return nil
}

func (s *SArticle) DeleteArticleByID(id uint64, ctx context.Context) error {
	_, err := g.MysqlDB.ExecContext(ctx, "delete from article where id=?", id)
	if err != nil {
		g.Logger.Error("delete mysql record failed.",
			zap.Error(err),
			zap.String("table", "article"))
		return fmt.Errorf("delete mysql record failed")
	}
	return nil

}

func (s *SArticle) UpdateArticleByID(id uint64, ctx context.Context) error {
	article := new(model.AnswerPost)
	_, err := g.MysqlDB.ExecContext(ctx, "update article set title=? ,content=?,update_time=? ",
		article.Title,
		article.Content,
		article.UpdateTime)
	if err != nil {
		g.Logger.Error("update mysql record failed.",
			zap.Error(err),
			zap.String("table", "article"))
		return fmt.Errorf("update mysal record failed")
	}
	return nil
}
