package model

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Canaries []*Canary `json:"canaries"`
	Status   string    //状态  未审核 已审核 驳回 正在执行 执行失败 执行成功
}

func (t *Task) SortCanariesByStage() {
	for i, canary1 := range t.Canaries {
		for j, canary2 := range t.Canaries {
			if j>i && canary1.Stage>canary2.Stage{
				t.Canaries[i],t.Canaries[j] =t.Canaries[j],t.Canaries[i]
			}
		}
	}
}
