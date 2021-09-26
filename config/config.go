package config

import (
	"os"
)

type Database struct {
	Dbtype   string `json:"dbtype"`
	Host     string `json:"host"`
	Port     int `json:"port"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

const (
	Dev = "dev"
	Test = "test"
	Prod = "prod"
)

var Mode =  Test

func init() {
	Mode = os.Getenv("RuntimeMode")
}
var DatabaseConfig Database

type Application struct {
	Host string
	Port string
}

var ApplicationConfig Application
func Setup(path string) {
	//1. 初始化 应用配置
	ApplicationConfig.Port="8002"
	ApplicationConfig.Host="127.0.0.1"
	//2. 初始化 数据库配置
	DatabaseConfig.Host="127.0.0.1"
	DatabaseConfig.Name="auto"
	DatabaseConfig.Dbtype="mysql"
	DatabaseConfig.Port=3306
	DatabaseConfig.Username="root"
	DatabaseConfig.Password="123456"
	//3. 初始化 外部服务配置
}