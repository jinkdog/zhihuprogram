package global

import (
	"database/sql"
	"github.com/go-redis/redis/v9"
	"go.uber.org/zap"
	"main/app/internal/model/config"
)

var ( //定义在global中的好处：其他包也能访问var里面的对象
	Config *config.Config
	//从app/internal/model包的config中导入Config结构体
	//config结构体又承接Logger，DataBase，App，Server，Middleware的结构体指针

	Logger *zap.Logger //从zap包导入//方便其他文件调用logger对象

	MysqlDB *sql.DB       //定义sql对象			//定义gormMysql对象
	Rdb     *redis.Client //定义redis对象
)
