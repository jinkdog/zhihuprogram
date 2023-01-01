package config

import (
	"time"
)

type Server struct {
	Mode         string `mapstructure:"mode" yaml:"mode"`
	Port         string `mapstructure:"port" yaml:"port"`
	ReadTimeout  string `mapstructure:"readTimeout" yaml:"readTimeout"`
	WriteTimeout string `mapstructure:"writeTimeout" yaml:"writeTimeout"`
}

func (s *Server) Addr() string { //sever结构体的方法
	return ":" + s.Port
}

func (s *Server) GetReadTimeout() time.Duration {
	t, _ := time.ParseDuration(s.ReadTimeout)
	//time.ParseDuration函数的作用是将string类型的时间转化为时间段类型
	return t
}

func (s *Server) GetWriteTimeout() time.Duration {
	t, _ := time.ParseDuration(s.WriteTimeout)
	return t
}
