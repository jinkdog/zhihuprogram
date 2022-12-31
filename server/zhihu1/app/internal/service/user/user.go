package user

import (
	"context" //上下文
	"encoding/hex"
	"errors"
	"fmt" //打印
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"golang.org/x/crypto/sha3"
	g "main/app/global"
	"main/app/internal/model"
)

type SUser struct{}

var insUser = SUser{} //自定义类型将insUser定义为SUser结构体

func (s *SUser) CheckUserIsExist(ctx context.Context, username string) error {
	userSubject := &model.UserSubject{}
	userSubject.Username = username

	_, err := g.MysqlDB.QueryContext(ctx, "select username from user_subject where username =?", userSubject.Username)

	//defer g.MysqlDB.Close()//这里需不需要关数据库？
	//后续这部分由问题

	if err != nil { //有错误
		var ErrRecordNotFound = errors.New("record not found")
		if err != ErrRecordNotFound { //错误不为已知的错误
			g.Logger.Error("query mysql record failed.",
				zap.Error(err),
				zap.String("table", "user_subject"),
			)
			return fmt.Errorf("internal err")
		}
	} else {
		return fmt.Errorf("username already exist")
	}

	//if err != nil {
	//	g.Logger.Error("query mysql record failed.",
	//		zap.Error(err),
	//		zap.String("table", "user_subject"))
	//}

	return nil
}

func (s *SUser) EncryptPassword(password string) string {
	d := sha3.Sum224([]byte(password))
	return hex.EncodeToString(d[:])
}

func (s *SUser) CreateUser(ctx context.Context, userSubject *model.UserSubject) {

	g.MysqlDB.ExecContext(ctx, "insert into user_subject(id, username, password, creattime, updatetime) values (?,?,?,?,?)",
		userSubject.Id,
		userSubject.Username,
		userSubject.Password,
		userSubject.CreateTime,
		userSubject.UpdateTime)
	//g.MysqlDB.Close()
	//g.MysqlDB.ExecContext(ctx, "insert into user_subject(username, password) values (?,?)",
	//	userSubject.Username,
	//	userSubject.Password)
	//g.MysqlDB.Close()
}

func (s *SUser) CheckPassword(ctx context.Context, userSubject *model.UserSubject) error {
	_, err := g.MysqlDB.QueryContext(ctx, "select password from user_subject where password=?", userSubject.Password)
	if err != nil {
		var ErrRecordNotFound = errors.New("record not found")
		if err != ErrRecordNotFound {
			g.Logger.Error("query mysql record failed.",
				zap.Error(err),
				zap.String("table", "user_subject"),
			)
			return fmt.Errorf("internal err")
		} else {
			return fmt.Errorf("invalid username or password")
		}
	}

	return nil
}

//
//func (s *SUser) GenerateToken(ctx context.Context, userSubject *model.UserSubject) (string, error) {
//	jwtConfig := g.Config.Middleware.Jwt
//
//	j := jwt.NewJWT(&jwt.Config{
//		SecretKey:   jwtConfig.SecretKey,
//		ExpiresTime: jwtConfig.ExpiresTime,
//		BufferTime:  jwtConfig.BufferTime,
//		Issuer:      jwtConfig.Issuer})
//	claims := j.CreateClaims(&jwt.BaseClaims{
//		Id:         userSubject.Id,
//		Username:   userSubject.Username,
//		CreateTime: userSubject.CreateTime,
//		UpdateTime: userSubject.UpdateTime,
//	})
//
//	tokenString, err := j.GenerateToken(&claims)
//	if err != nil {
//		g.Logger.Error("generate token failed.", zap.Error(err))
//		return "", fmt.Errorf("internal err")
//	}
//
//	err = g.Rdb.Set(ctx,
//		fmt.Sprintf("jwt:%d", userSubject.Id),
//		tokenString,
//		time.Duration(jwtConfig.ExpiresTime)*time.Second).Err()
//	if err != nil {
//		g.Logger.Error("set redis cache failed.",
//			zap.Error(err),
//			zap.String("key", "jwt:[id]"),
//			zap.Int64("id", userSubject.Id),
//		)
//		return "", fmt.Errorf("internal err")
//	}
//
//	return tokenString, nil
//}
