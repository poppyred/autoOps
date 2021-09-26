package model

import "gorm.io/gorm"

//内置脚本
type Script struct {
	gorm.Model
	ScriptType string
	ScriptBody string
	Privilege  string
	//出参
	//入参
	//权限
}
