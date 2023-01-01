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
	"main/utils/jwt"
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

func (s *SUser) GenerateToken(ctx context.Context, userSubject *model.UserSubject) (string, error) { //生成Token

	jwtConfig := g.Config.Middleware.Jwt
	//对应的是model/config/config.go里的middleware结构体
	//因为model/config目录下的结构体都被集成到了config.go中
	//生成的是jwt的配置文件

	j := jwt.NewJWT(&jwt.Config{ //NewJWT函数接受jwt配置文件内存的地址
		SecretKey:   jwtConfig.SecretKey,
		ExpiresTime: jwtConfig.ExpiresTime,
		BufferTime:  jwtConfig.BufferTime,
		Issuer:      jwtConfig.Issuer})
	claims := j.CreateClaims(&jwt.BaseClaims{
		//CreateClaims是*Jwt.jwt的一个方法输入基本声明，返回自定义声明CustomClaims
		//该方法包装在utils层中
		Id:         userSubject.Id,
		Username:   userSubject.Username,
		CreateTime: userSubject.CreateTime,
		UpdateTime: userSubject.UpdateTime,
	})
	//CustomClaims的自定义类型返回的是一个
	//type CustomClaims struct {
	//	BufferTime int64//缓冲时间
	//	jwt.RegisteredClaims
	//	BaseClaims//基本声明
	//}
	//type RegisteredClaims struct {//注册声明
	//	Issuer    string       `json:"iss,omitempty"`//发行人
	//	Subject   string       `json:"sub,omitempty"`//主题
	//	Audience  ClaimStrings `json:"aud,omitempty"`//授予对象
	//	ExpiresAt *NumericDate `json:"exp,omitempty"`//过期时间
	//	NotBefore *NumericDate `json:"nbf,omitempty"`//token生效时间
	//	IssuedAt  *NumericDate `json:"iat,omitempty"`//签发时间
	//	ID        string       `json:"jti,omitempty"`
	//}
	//但是实际上的CustomClaims中的RegisteredClaims并没有声明这么多
	//详见zhihu1/utils/jwt/jwt.go

	tokenString, err := j.GenerateToken(&claims)
	if err != nil {
		g.Logger.Error("generate token failed.", zap.Error(err))
		return "", fmt.Errorf("internal err")
	} //生成token失败，返回一个空字段和错误

	err = g.Rdb.Set(ctx, //上下文
		fmt.Sprintf("jwt:%d", userSubject.Id),            //键
		tokenString,                                      //值
		time.Duration(jwtConfig.ExpiresTime)*time.Second, //持续时间
	).Err()
	//Rdb为*redis.Client，而Client结构体中包含cmdable
	//cmdable又可以调用Set函数返回*StatusCmd
	//StatusCmd中有包含baseCmd可以调用Err（）返回一个错误
	if err != nil {
		g.Logger.Error("set redis cache failed.", //设置redis缓存失败
			zap.Error(err),
			zap.String("key", "jwt:[id]"),
			zap.Int64("id", userSubject.Id),
		)
		return "", fmt.Errorf("internal err")
	} //返回一个空的token和错误

	//如果成功生成tokenString以及在redis中设置集合来储存tokenString也成功那么就可以成功返回
	return tokenString, nil
}
