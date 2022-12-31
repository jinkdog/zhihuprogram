package user

import (
	"context" //上下文
	"encoding/hex"
	"fmt" //打印
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"golang.org/x/crypto/sha3"
	g "main/app/global"
	"main/app/internal/model"
	"time"
)

type SUser struct{}

var insUser = SUser{} //自定义类型将insUser定义为SUser结构体

func (s *SUser) CheckUserIsExist(ctx context.Context, username string) error {
	userSubject := &model.UserSubject{}

	rows, err := g.MysqlDB.QueryContext(ctx, "select * from user_subject where username =?", username)
	//defer g.MysqlDB.Close()//这里需不需要关数据库？
	//后续这部分由问题
	//每次这里会报错，说明存在错误所以是QueryContext这个函数的搜索结果返回了错误
	//但是QueryContext由于第一次注册时没有存名字，所以如果没有找到数据，他就会返回错误，而这个错误会在db。Scan中被调用，最终返回一个error
	for rows.Next() {

		if err != nil { //有错误
			if err != rows.Scan(&userSubject.Username) {
				g.Logger.Error("query mysql record failed.",
					zap.Error(err),
					zap.String("table", "user_subject"),
				)
				return fmt.Errorf("internal err")
			}
		} else {
			return fmt.Errorf("username already exist")
		}
	}

	return nil
}

func (s *SUser) EncryptPassword(password string) string {
	d := sha3.Sum224([]byte(password))
	return hex.EncodeToString(d[:])
}

func (s *SUser) CreateUser(ctx context.Context, userSubject *model.UserSubject) {

	_, err := g.MysqlDB.ExecContext(ctx, "insert into user_subject(id, username, password, creattime, updatetime) values (?,?,?,?,?)",
		userSubject.Id,
		userSubject.Username,
		userSubject.Password,
		time.Now(),
		time.Now())
	if err != nil {
		g.Logger.Error("query mysql record failed.",
			zap.Error(err),
			zap.String("table", "user_subject"),
		)

	}

	//g.MysqlDB.ExecContext(ctx, "insert into user_subject(username, password) values (?,?)",
	//	userSubject.Username,
	//	userSubject.Password)
	//g.MysqlDB.Close()
}

func (s *SUser) CheckPassword(ctx context.Context, userSubject *model.UserSubject) error {
	rows, err := g.MysqlDB.QueryContext(ctx, "select password from user_subject where password=?", userSubject.Password)
	for rows.Next() {
		if err != nil {
			if err != rows.Scan(&userSubject.Password) {
				g.Logger.Error("query mysql record failed.",
					zap.Error(err),
					zap.String("table", "user_subject"),
				)
				return fmt.Errorf("internal err")
			} else {
				return fmt.Errorf("invalid username or password")
			}
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
