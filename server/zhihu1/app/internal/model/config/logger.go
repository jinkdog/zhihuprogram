package config

type Logger struct {
	SavePath   string `mapstructure:"savePath" yaml:"savePath"`
	MaxSize    int    `mapstructure:"maxSize" yaml:"maxSize"`
	MaxAge     int    `mapstructure:"maxAge" yaml:"maxAge"`
	MaxBackups int    `mapstructure:"maxBackups" yaml:"maxBackups"`
	IsCompress bool   `mapstructure:"isCompress" yaml:"isCompress"`
	LogLevel   string `mapstructure:"logLevel" yaml:"logLevel"`
}

//加tag原因：
//1与yaml文件中的信息一一对应，yaml文件中savePath是什么，yaml：什么
//2在viper中的反序列化能找到对应的对象

//yaml中的文件一一对应
//配置文件既要写在结构体里面，也要写在yaml文件里面

//数据类型也要和getWriteSyncer函数返回的类型一一对应
