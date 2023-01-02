package config

type Middleware struct { //中间层结构体包括CORS指针和Jwt指针
	Cors *CORS `mapstructure:"cors" yaml:"cors"`
	Jwt  *Jwt  `mapstructure:"jwt" yaml:"jwt"`
}

type CORS struct {
	Mode      string          `mapstructure:"mode" yaml:"mode"`
	Whitelist []CORSWhitelist `mapstructure:"whitelist" yaml:"whitelist"`
}

type CORSWhitelist struct {
	AllowOrigin      string `mapstructure:"allowOrigin" yaml:"allowOrigin"`
	AllowMethods     string `mapstructure:"allowMethods" yaml:"allowMethods"`         //允许的方法
	AllowHeaders     string `mapstructure:"allowHeaders" yaml:"allowHeaders"`         //允许的头文件
	ExposeHeaders    string `mapstructure:"exposeHeaders" yaml:"exposeHeaders"`       //暴露的头文件
	AllowCredentials bool   `mapstructure:"allowCredentials" yaml:"allowCredentials"` //允许的证书
}

type Jwt struct {
	SecretKey   string `mapstructure:"secretKey" yaml:"secretKey"`     //密钥
	ExpiresTime int64  `mapstructure:"expiresTime" yaml:"expiresTime"` //过期时间
	BufferTime  int64  `mapstructure:"bufferTime" yaml:"bufferTime"`   //缓冲时间
	Issuer      string `mapstructure:"issuer" yaml:"issuer"`           //发行人
}
