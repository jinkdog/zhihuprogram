package config

type Config struct { //与viper相关
	Logger *Logger `mapstructure:"logger" yaml:"logger"`

	DataBase *Database `mapstructure:"database"  yaml:"database"`

	App *App `mapstructure:"app"  yaml:"app"`

	Server *Server `mapstructure:"server"  yaml:"server"`

	Middleware *Middleware `mapstructure:"middleware" yaml:"middleware"`
}

//传递的都为指针，使得运算效率加快
//config接受app，database，logger，middleware，server的指针可以调用数据，使用其对应的方法
//这几个的配置文件信息都在yaml里所以要映射到config文件夹中
