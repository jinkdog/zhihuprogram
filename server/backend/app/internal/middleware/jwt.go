package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	g "main/app/global"
	"main/app/internal/model/response"
	"main/utils/cookie"
	myjwt "main/utils/jwt"
	"net/http"
	"time"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	//该函数返回的就是gin的中间件
	return func(c *gin.Context) {
		var token string

		cookieConfig := g.Config.App.Cookie //获取配置文件里的Cookie的配置信息

		cookieWriter := cookie.NewCookieWriter(cookieConfig.Secret,
			cookie.Option{
				Config: cookieConfig.Cookie,
				Ctx:    c,
			}) //好像把utils包里的cookie和internal\config包里的cookie联系起来了
		//将config包里的cookie信息转化为utils包里的cookie信息

		ok := cookieWriter.Get("x-token", &token)
		if token == "" || !ok { //utils包里的cookie包中的Get方法
			response.Fail(c, http.StatusUnauthorized, 1, "not logged in")
			c.Abort() //表示只有该中间件被停止运行，其他中间件正常运行
			return
		}

		jwtConfig := g.Config.Middleware.Jwt //获取配置文件里Jwt的配置信息
		j := myjwt.NewJWT(&myjwt.Config{
			//"main/utils/jwt"这里将这个包重命名为myjwt以免和其他代码出现虽然只有名字是一样的包引入
			//引入的config文件来自util包里的结构体
			SecretKey: jwtConfig.SecretKey,
		})

		mc, err := j.ParseToken(token) //解析token//返回一个自定义声明
		if err != nil {
			response.Fail(c, http.StatusBadRequest, 1, err.Error())
			c.Abort() //如果发生错误，把其他中间件隔离在外继续运行，该中间件不运行
			return
		}

		if mc.ExpiresAt.Unix()-time.Now().Unix() < mc.BufferTime && mc.ExpiresAt.Unix()-time.Now().Unix() > 0 {
			//如果过期的时间转换为距离1970...的int数减去现在到1970...的int数小于缓存的时间
			//并且过期的时间转换为距离1970...的int数减去现在到1970...的int数小于缓存的时间大于0
			mc.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Duration(g.Config.Middleware.Jwt.ExpiresTime) * time.Second))

			//现在的时间加上yaml文件里配置的时间（用time.Duration进行强制的类型转换为int64）乘以单位时间秒生成一个过期的时间点
			newToken, _ := j.GenerateToken(mc)     //根据自定义声明返回一个tokenstring
			newClaims, _ := j.ParseToken(newToken) //解析tokenstring返回一个新声明
			cookieWriter.Set("x-token", newToken)

			err = g.Rdb.Set(c,
				fmt.Sprintf("jwt:%d", newClaims.BaseClaims.Id),
				newToken,
				time.Duration(jwtConfig.ExpiresTime)*time.Second).Err()
			//传入gin的上下文，新声明中基本信息的id，新的token，过期的时间，最后得到的*StatusCmd里面的 baseCmd调用的输出错误的方法

			if err != nil {
				g.Logger.Error("set redis key failed.",
					zap.Error(err),
					zap.String("key", "jwt:[id]"), zap.Int64("id", newClaims.BaseClaims.Id),
				)
				response.InternalErr(c)

				c.Abort()
				return
			}
		}

		c.Set("id", mc.BaseClaims.Id)
		c.Set("username", mc.BaseClaims.Username)
		c.Next()
	}
}
