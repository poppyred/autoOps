package migrate

import (
	database "autoOps/pkg/db"
	config "autoOps/config"
	"autoOps/pkg/service"
	"fmt"
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
)

var (
	configPath  string
	mode     string
	StartCmd = &cobra.Command{
		Use:   "init",
		Short: "initialize the database",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {

	StartCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "settings.dev.yml", "Start server with provided configuration file")
	StartCmd.PersistentFlags().StringVarP(&mode, "mode", "m", "dev", "server mode ; eg:dev,test,prod")
}

func run() {
	usage := `start init`
	config.Mode=mode
	fmt.Println(usage)
	//1. 读取配置
	config.Setup(configPath)
	//2. 初始化数据库链接
	database.Setup()
	//3. 数据库迁移
	_ = migrateModel()
	klog.Info("数据库结构初始化成功！")
	klog.Info(usage)
}

func migrateModel() error {
	if config.DatabaseConfig.Dbtype == "mysql" {
		database.MysqlDB =database.MysqlDB.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4")
	}
	return service.Migrate(database.MysqlDB)
}