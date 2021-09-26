package service

import (
	"autoOps/pkg/model"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	//db.SingularTable(true)
	return db.AutoMigrate(
		// 初始化表
		new(model.Agent),
		new(model.Script),
		new(model.TaskResult),
		new(model.Canary),
		new(model.Task),
	)
}
