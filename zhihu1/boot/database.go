package boot

import (
	"database/sql"
	"fmt"
	"github.com/go-redis/redis/v9"
	"go.uber.org/zap" //日志
	"golang.org/x/net/context"
	g "main/app/global"
	"time"
)

func MysqlDBSetup() {
	config := g.Config.DataBase.Mysql //利用config将Database结构体引入

	db, err := sql.Open("mysql", config.GetDsn())
	if err != nil {
		g.Logger.Fatal("initialize mysql failed.", zap.Error(err))
		//fatal的作用：打印日志之后整个项目直接退出
	}

	sqlDB := db
	sqlDB.SetConnMaxIdleTime(g.Config.DataBase.Mysql.GetConnMaxIdleTime()) //传入时间段
	sqlDB.SetConnMaxLifetime(g.Config.DataBase.Mysql.GetConnMaxLifeTime()) //传入时间段
	sqlDB.SetMaxIdleConns(g.Config.DataBase.Mysql.MaxIdleConns)
	sqlDB.SetMaxOpenConns(g.Config.DataBase.Mysql.MaxOpenConns)
	err = sqlDB.Ping() //测试数据是否连接成功
	if err != nil {
		g.Logger.Fatal("connect to mysql db failed.", zap.Error(err)) //连接失败则使用fatal函数
	}

	g.MysqlDB = db //gorm的db复制到global中的MySQL里面

	g.Logger.Info("initialize mysql successfully!")
}

func RedisSetup() {
	config := g.Config.DataBase.Redis

	rdb := redis.NewClient(&redis.Options{ //定义新的客户端
		Addr:     fmt.Sprintf("%s:%s", config.Addr, config.Port), //将config中的地址和端口连接起来构成字符串
		Username: config.Username,                                //config中的用户名
		Password: config.Password,                                //config中的密码
		DB:       config.Db,                                      //config中的数据库名
		PoolSize: config.PoolSize,                                //config中的连接池名
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) //测试连接是否成功
	defer cancel()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		g.Logger.Fatal("connect to redis instance failed.", zap.Error(err))
	}

	g.Rdb = rdb

	g.Logger.Info("initialize redis client successfully!")
}
