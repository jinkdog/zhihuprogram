package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type JWT struct { //JWT包装了config结构体的指针
	Config *Config
}

type Config struct {
	SecretKey   string // 密钥
	ExpiresTime int64  // 过期时间,单位:秒
	BufferTime  int64  // 缓冲时间,缓冲时间内会获得新的token刷新令牌,此时一个用户会存在两个有效令牌,但是前端只留一个,另一个会丢失
	Issuer      string // 签发者
}

var (
	TokenExpired     = errors.New("token is expired")           //令牌过期
	TokenNotValidYet = errors.New("token not active yet")       //令牌无效
	TokenMalformed   = errors.New("that's not even a token")    //没有令牌
	TokenInvalid     = errors.New("couldn't handle this token") //无法处理令牌
)

func NewJWT(config *Config) *JWT { //生成JWT
	return &JWT{Config: config}
}

func (j *JWT) CreateClaims(baseClaims *BaseClaims) CustomClaims { //是JWT结构体的方法
	claims := CustomClaims{
		BufferTime: j.Config.BufferTime, //缓冲时间
		RegisteredClaims: jwt.RegisteredClaims{ //注册声明
			NotBefore: jwt.NewNumericDate(time.Now().Truncate(time.Second)), // 签名生效时间
			//先通过time.Now（）返回一个现在的时间time
			//然后再Time.Truncate（）根据传入的单位时间长度来约分成时间戳（一个时间点）
			//jwt.NewNumericDate通过原来的时间戳，返回一个JSON格式的时间戳
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(j.Config.ExpiresTime) * time.Second)),
			//现在的时间加上配置文件中的过期时间段
			//time.Duration（）是将配置文件中的ExpiresTime的int64类型强制转换为实质上也是int64但叫做Duration的类型
			//这个类型可以和time.Second作乘法获得时间的单位，最终形成一个时间段
			//最后再将这个时间戳转换为JSON的时间戳
			Issuer: j.Config.Issuer,
		},
		BaseClaims: *baseClaims,
	}
	return claims
}

func (j *JWT) GenerateToken(claims *CustomClaims) (string, error) {
	//仍是JWT的一个方法，传入自定义声明，返回string，error值
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, *claims)
	//通过传入一个自定义的声明和选择的编码方式来输出一个token
	signingKey := []byte(j.Config.SecretKey) //将string类型切割成byte类型，并储存在byte数组
	return token.SignedString(signingKey)    //根据传入的一个任意值来获得一个string（JWT）和error
}

// ParseToken 解析JWT
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	// 解析token，传入tokenString和自定义声明
	signingKey := []byte(j.Config.SecretKey) //传入密钥并将其转换为字符数组
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{},
		func(token *jwt.Token) (i interface{}, err error) { //传入一个解析的匿名函数
			return signingKey, nil
		})
	//传入tokenString、CustomClaims（自定义声明，但是一开始传入ParseToken的是CustomClaims的指针所以要取地址）
	//还要传入Keyfunc//options ...ParserOption（这个是提供额外解析配置的选项，可以不添加）
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			//此处使用了go语言中的结构体类型断言
			//err是一个包含Error（）函数的接口
			//jwt.ValidationError则是一个结构体包含Inner error，Errors uint32，text string的成员，以及相同的Error（）的方法
			//因而可以实现类型断言
			//如果断言成功则返回ok为true，如果断言失败则返回ok为false
			if ve.Errors&jwt.ValidationErrorMalformed != 0 { //误操作的验证错误
				//如果ve.Errors和jwt.ValidationErrorMalformed都不为0
				return nil, TokenMalformed
				//var TokenMalformed error = errors.New("that's not even a token")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 { //验证过期的错误
				// Token is expired
				return nil, TokenExpired
				//var TokenExpired error = errors.New("token is expired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 { //验证无效的错误
				return nil, TokenNotValidYet
				//var TokenNotValidYet error = errors.New("token not active yet")

			} else {
				return nil, TokenInvalid
				//var TokenInvalid error = errors.New("couldn't handle this token")
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			//再次做一个类型断言
			//token.Claims是一个包含Valid() error的接口
			//CustomClaims结构体中的 jwt.RegisteredClaims可以实现 Valid() error方法
			//因此可以实现类型断言
			//不过这是怎么想出来的？
			//如果断言成功，而且token有效
			return claims, nil
		}
		return nil, TokenInvalid
	} else { //如果token为空
		return nil, TokenInvalid
	}
}
