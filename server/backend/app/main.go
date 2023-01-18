package main

import "main/boot"

func main() {
	boot.ViperSetup()   //运行viper组件
	boot.LoggerSetup()  //运行日志组件
	boot.MysqlDBSetup() //运行MySQL组件
	boot.RedisSetup()   //运行redis组件
	boot.ServerSetup()  //运行gin框架服务
}
