package boot

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	g "main/app/global"
	"main/app/router"
	"net/http"
)

func ServerSetup() {
	config := g.Config.Server //将sever结构体导入

	gin.SetMode(config.Mode) //将config对应的string输入//具体值在yaml文件当中
	//debug模式用于调试
	//release模式将项目部署到服务器中
	routers := router.InitRouter()
	server := &http.Server{
		Addr:              config.Addr(),
		Handler:           routers,
		TLSConfig:         nil,
		ReadTimeout:       config.GetReadTimeout(),
		ReadHeaderTimeout: 0,
		WriteTimeout:      config.GetWriteTimeout(),
		IdleTimeout:       0,
		MaxHeaderBytes:    1 << 20, // 16 MB
	}

	g.Logger.Info("initialize server successfully!", zap.String("port", config.Addr()))
	g.Logger.Error(server.ListenAndServe().Error())
}

//大概的作用是通过引用router中的函数打开路由
