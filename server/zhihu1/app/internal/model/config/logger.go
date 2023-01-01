package config

type Logger struct {
	SavePath   string `mapstructure:"savePath" yaml:"savePath"`     //日志保存路径
	MaxSize    int    `mapstructure:"maxSize" yaml:"maxSize"`       //日志最大大小
	MaxAge     int    `mapstructure:"maxAge" yaml:"maxAge"`         //日志最长保存时间为
	MaxBackups int    `mapstructure:"maxBackups" yaml:"maxBackups"` //最大备份的份数
	IsCompress bool   `mapstructure:"isCompress" yaml:"isCompress"` //是否压缩
	LogLevel   string `mapstructure:"logLevel" yaml:"logLevel"`     //日志等级
}

//加tag原因：
//1与yaml文件中的信息一一对应，yaml文件中savePath是什么，yaml：什么
//2在viper中的反序列化能找到对应的对象

//yaml中的文件一一对应
//配置文件既要写在结构体里面，也要写在yaml文件里面

//数据类型也要和getWriteSyncer函数返回的类型一一对应
