package middleware

import (
	"github.com/gin-gonic/gin"
	g "main/app/global"
	"main/app/internal/model/config"
	"net/http"
)

func checkCors(currentOrigin string) *config.CORSWhitelist {
	for _, whitelist := range g.Config.Middleware.Cors.Whitelist {
		// 遍历配置中的cors头并进行匹配
		if currentOrigin == whitelist.AllowOrigin {
			//如果传入的现在的源和白名单的源是一样的那么就可以返回一个遍历后的
			return &whitelist
		}
	}
	return nil
}

// Cors allow all cors request
func Cors() gin.HandlerFunc { //Cors中间件
	return func(c *gin.Context) {
		method := c.Request.Method               //承接网页请求的方式如GET、POST等
		origin := c.Request.Header.Get("Origin") //获得源
		c.Header("Access-Control-Allow-Origin", origin)
		//获取跨域可访问的源
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,X-Token,X-Sign-Id")
		//获取跨域可访问的头文件
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS,DELETE,PUT")
		//获取跨域可访问的方式
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		//获取跨域访问中暴露的文件信息
		c.Header("Access-Control-Allow-Credentials", "true")
		//获取跨域访问的许可证

		// allow all method
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		} //如果请求方法为OPTIONS说明没有要传递的内容，这个时候用AbortWithStatus相当于停止运行该中间件并返回状态码

		// 处理请求
		c.Next() //继续运行接下来的中间件
	}
}

// CorsByRules process request base on configured logic
func CorsByRules() gin.HandlerFunc {
	// allow all
	if g.Config.Middleware.Cors.Mode == "allow-all" {
		//当配置文件中的Cors状态为"allow-all"使用Cors（）函数
		return Cors()
	}
	return func(c *gin.Context) { //如果不是"allow-all"那么就返回下面这个中间件
		whitelist := checkCors(c.GetHeader("origin"))

		// passed, add request header
		if whitelist != nil { //成功获得到白名单，而且获得证书允许
			c.Header("Access-Control-Allow-Origin", whitelist.AllowOrigin)
			c.Header("Access-Control-Allow-Headers", whitelist.AllowHeaders)
			c.Header("Access-Control-Allow-Methods", whitelist.AllowMethods)
			c.Header("Access-Control-Expose-Headers", whitelist.ExposeHeaders)
			if whitelist.AllowCredentials {
				c.Header("Access-Control-Allow-Credentials", "true")
			}
		}

		// not passed, deny
		if whitelist == nil && g.Config.Middleware.Cors.Mode == "strict-whitelist" && !(c.Request.Method == "GET" && c.Request.URL.Path == "/health") {
			//如果白名单为空
			//而且配置文件中的模式为"strict-whitelist"
			//而且请求方式不为GET，请求路径不再health下
			c.AbortWithStatus(http.StatusForbidden)
			//停止运行该中间件并返回禁止状态码
		} else {
			// allow all method no matter it passed or not
			if c.Request.Method == "OPTIONS" {
				//当请求方式为OPTION时同样停止运行该中间件并返回没有内容的状态码
				c.AbortWithStatus(http.StatusNoContent)
			}
		}

		c.Next()
	}
}
