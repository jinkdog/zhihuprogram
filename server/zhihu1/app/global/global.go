package global

import (
	"database/sql"
	"go.uber.org/zap"
	"main/app/internal/model/config"
)

var ( //定义在global中的好处：其他包也能访问var里面的对象
	Config *config.Config //从internal包的config中导入Config结构体

	Logger *zap.Logger //从zap包导入//方便其他文件调用logger对象

	MysqlDB *sql.DB //定义sql对象//定义gormMysql对象
	//Rdb     *redis.Client //定义redis对象
)
