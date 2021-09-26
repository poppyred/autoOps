package model

import (
	database "autoOps/pkg/db"
	"autoOps/pkg/types"
	"gorm.io/gorm"
	"time"
)

type Agent struct {
	gorm.Model
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Addr        string    `json:"addr"`
	LastPing    time.Time `json:"lastPing"`
	Status      string    `json:"status"` //注册过来默认禁用
	Canaries    []Canary `gorm:"many2many:agent_tasks;" json:"canaries"`
}

func NewAgent() *Agent {
	return &Agent{}
}

func (a *Agent) Create() error {
	return database.MysqlDB.Create(a).Error
}

func (a *Agent) Delete() error {
	return database.MysqlDB.Delete(a, a.ID).Error
}

func (a *Agent) Update() error {
	return database.MysqlDB.Delete(a, a.ID).Error
}

func (a *Agent) FindByID() error {
	return database.MysqlDB.Find(a, a.ID).Error
}

func (a *Agent) List(q *types.Pagination) (agents []Agent, err error) {
	err = database.MysqlDB.Limit(q.PageSize).Offset((q.Page - 1) * q.PageSize).Find(&agents).Error
	return
}
