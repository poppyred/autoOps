package model

import "gorm.io/gorm"

type TaskResult struct {
	gorm.Model
	Stdout   string `json:"stdout"`
	Stderr   string `json:"stderr"`
	OutParam string `json:"outParam"`
	CanaryID uint   `json:"canaryId"`
	Canary   Canary `json:"canary"`
	AgentID  uint   `json:"agentId"`
	Agent    Agent  `json:"agent"`
}
