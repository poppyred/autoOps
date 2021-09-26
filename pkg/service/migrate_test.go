package service

import (
	"autoOps/config"
	database "autoOps/pkg/db"
	"testing"
)

func TestMigrate(t *testing.T) {
	config.Setup("")
	database.Setup()
	err := Migrate(database.MysqlDB)
	if err != nil {
		t.Error(err)
	}
}
