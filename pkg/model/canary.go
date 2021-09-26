package model

import "gorm.io/gorm"

//金丝雀
type Canary struct {
	gorm.Model
	Stage  int `json:"stage"` //灰度步骤
	Task   Task
	TaskID uint `json:"taskId"`
	//Script Script 停留执行时的脚本 todo 还是否需要更新脚本
	IsBuildIn    *bool     `json:"isBuildIn"`  //内置则为用户上传
	Script       *string   `json:"script"`     //用户上传则为nil
	ScriptType   *string   `json:"scriptType"` //用户上传则为nil
	TargetAgents []*Agent  `gorm:"many2many:agent_tasks;" json:"targetAgents"`
	NextCanary   []*Canary `gorm:"many2many:next_canary;" json:"nextCanary"` //todo 下一期预留
}
