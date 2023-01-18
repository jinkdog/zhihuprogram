package comment

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	g "main/app/global"
	"main/app/internal/model"
	"time"
)

type SComment struct{}

var insComment = SComment{}

func (s *SComment) CreateCommend(ctx context.Context) error {
	comment := new(model.Comment)
	_, err := g.MysqlDB.ExecContext(ctx, "insert into comment(comment_id,content,post_id,author_id,parent_id,create_time,update_time)values(?,?,?,?,?,?,?,?)",
		comment.CommentId,
		comment.Content,
		comment.PostId,
		comment.AuthorId,
		comment.ParentId,
		time.Now(),
		time.Now())

	if err != nil {
		g.Logger.Error("insert mysql record failed.",
			zap.Error(err),
			zap.String("table", "comment"))
		return fmt.Errorf("insert mysql record failed")
	}
	return nil

}

func (s *SComment) GetCommentListById(id uint64, ctx context.Context) error {
	comment := new(model.Comment)
	rows, err := g.MysqlDB.QueryContext(ctx, "select comment_id,content,post_id,author_id,parent_id,create_time from comment where comment_id=?", id)

	for rows.Next() {
		if err != nil {
			if err != rows.Scan(&comment.CommentId, &comment.Content, &comment.PostId, &comment.AuthorId, &comment.ParentId) {
				g.Logger.Error("query mysql record failed.",
					zap.Error(err),
					zap.String("table", "comment"),
				)
				return fmt.Errorf("internal err")
			} else {
				return fmt.Errorf("no comment in the db")
			}
		}
	}
	return nil
}
