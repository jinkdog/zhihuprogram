package config

import (
	"fmt"
	"time"
)

type Database struct {
	//使用指针传递的好处
	//只用传第一个指针而不用将所有的数据传递
	Mysql *Mysql `mapstructure:"mysql" yaml:"mysql"`

	Redis *Redis `mapstructure:"redis" yaml:"redis"`
}

type Mysql struct { //与yaml文件中的一一对应
	Addr     string `mapstructure:"addr" yaml:"addr"`
	Port     string `mapstructure:"port" yaml:"port"`
	Db       string `mapstructure:"db" yaml:"db"`
	Username string `mapstructure:"username" yaml:"username"`
	Password string `mapstructure:"password" yaml:"password"`
	Charset  string `mapstructure:"charset" yaml:"charset"`

	ConnMaxIdleTime string `mapstructure:"connMaxIdleTime" yaml:"connMaxIdleTime"`
	ConnMaxLifeTime string `mapstructure:"connMaxLifeTime" yaml:"connMaxLifeTime"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns" yaml:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns" yaml:"maxOpenConns"`
}

func (m *Mysql) GetDsn() string { //将所有配置的字符串连接起来，形成一个dsn//结构体方法也要传入指针
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Australia%%2FMelbourne", //拼接字符串
		m.Username,
		m.Password,
		m.Addr,
		m.Port,
		m.Db,
		m.Charset)
}

func (m *Mysql) GetConnMaxIdleTime() time.Duration {
	//boot包中SetConnMaxIdleTime函数传入的时间段，此处将时间点转化为时间段
	t, _ := time.ParseDuration(m.ConnMaxIdleTime)
	return t
}

func (m *Mysql) GetConnMaxLifeTime() time.Duration { //此处同上
	t, _ := time.ParseDuration(m.ConnMaxLifeTime)
	return t
}

type Redis struct {
	Addr     string `mapstructure:"addr" yaml:"addr"`         //地址
	Port     string `mapstructure:"port" yaml:"port"`         //端口
	Username string `mapstructure:"username" yaml:"username"` //用户名
	Password string `mapstructure:"password" yaml:"password"` //密码
	Db       int    `mapstructure:"db" yaml:"db"`             //数据库名
	PoolSize int    `mapstructure:"poolSize" yaml:"poolSize"` //连接池大小
}
