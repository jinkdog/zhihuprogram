package user

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	g "main/app/global"
	"main/app/internal/model"
	"time"
)

type SCollect struct{}

var insCollect = SCollect{}

func (s *SCollect) CheckCollectionIsExist(ctx context.Context, collectType int32, userId int64, id interface{}) error {
	whereSql := ""
	switch collectType {
	case 1:
		whereSql = fmt.Sprintf("user_id = ? , restaurant_id = ?")
	case 2:
		whereSql = fmt.Sprintf("user_id = ? , recipe_id = ?")
	default:

	}

	userCollection := &model.UserCollection{}
	rows, err := g.MysqlDB.QueryContext(ctx, "select *from user_collection where "+whereSql, userId, id)
	for rows.Next() {
		if err != nil {
			if err != rows.Scan(&userCollection) {
				g.Logger.Error("query [user_collection] record failed",
					zap.Error(err),
					zap.String("table", "user_collection"))
				return fmt.Errorf("internal err")
			}
		} else {
			return fmt.Errorf("duplicate collect")
		}
	}

	return nil
}

func (s *SCollect) CheckCollectionIdIsExist(ctx context.Context, id, userId int64) error {
	userCollection := &model.UserCollection{}
	rows, err := g.MysqlDB.QueryContext(ctx, "select id,user_id from user_collection where id=?,user_id=?", id, userId)
	for rows.Next() {
		if err != nil {
			if err != rows.Scan(&userCollection) {
				g.Logger.Error("query mysql record failed",
					zap.Error(err),
					zap.String("table", "user_collection"))
				return fmt.Errorf("internal error")
			} else {
				return fmt.Errorf("collection not found")
			}

		}
	}

	return nil
}

func (s *SCollect) CreateCollection(ctx context.Context) error {
	userCollection := new(model.UserCollection)
	_, err := g.MysqlDB.ExecContext(ctx, "insert into user_collection(id,user_id,collect_type,question_id,answer_id,create_time,update_time)values(?,?,?,?,?,?,?)",
		userCollection.Id,
		userCollection.UserId,
		userCollection.CollectType,
		userCollection.QuestionId,
		userCollection.AnswerId,
		time.Now(),
		time.Now())
	if err != nil {
		g.Logger.Error("insert mysql record failed.",
			zap.Error(err),
			zap.String("table", "user_collection"),
		)
	}
	return nil
}
func (s *SCollect) DeleteCollectionByID(ctx context.Context, id int64) error {
	_, err := g.MysqlDB.ExecContext(ctx, "delete from user_collection where id=?", id)
	if err != nil {
		g.Logger.Error("delete mysql record failed.",
			zap.Error(err),
			zap.String("table", "question_community"))
		return fmt.Errorf("delete mysql record failed")
	}
	return nil
}

//func (s *sCollect) GetUserCollectionCount(ctx context.Context, userId int64, collectType int32) (int64, error) {
//	var cnt int64
//	rows, err := g.MysqlDB.QueryContext(ctx, "select *from user_collection where user_id=?collection_type=?", userId)
//	for rows.Next() {
//
//	}
//	err := g.MysqlDB.WithContext(ctx).
//		Table("user_collection").
//		Where("user_id = ? AND collect_type = ?", userId, collectType).
//		Count(&cnt).Error
//	if err != nil {
//		g.Logger.Errorf("query [user_collection] record failed ,err: %v", err)
//		return -1, fmt.Errorf("internal err")
//	}
//
//	return cnt, nil
//}
//
//func (s *sCollect) GetUserCollectionsWithLimit(ctx context.Context, userId int64, collectType int32, limit, page int) ([]*model.UserCollection, error) {
//	var userCollections []*model.UserCollection
//	err := g.MysqlDB.WithContext(ctx).
//		Table("user_collection").
//		Limit(limit).Offset(limit*(page-1)).
//		Where("user_id = ? AND collect_type = ?", userId, collectType).
//		Find(&userCollections).Error
//	if err != nil {
//		g.Logger.Errorf("query [user_collection] failed, err: %v", err)
//		return nil, fmt.Errorf("internal err")
//	}
//
//	return userCollections, nil
//}
