package server

import (
	"autoOps/config"
	database "autoOps/pkg/db"
	"autoOps/pkg/router"
	"autoOps/tools"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
	configPath string
	port       string
	mode       string
	StartCmd   = &cobra.Command{
		Use:     "server",
		Short:   "Start API server",
		Example: "me server settings.dev.yml",
		PreRun: func(cmd *cobra.Command, args []string) {
			usage()
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "settings.dev.yml", "Start server with provided configuration file")
	StartCmd.PersistentFlags().StringVarP(&port, "port", "p", "8002", "Tcp port server listening on")
	StartCmd.PersistentFlags().StringVarP(&mode, "mode", "m", "dev", "server mode ; eg:dev,test,prod")
}

func usage() {
	usageStr := `starting api server`
	log.Printf("%s\n", usageStr)
}

func setup() {

	// 1. 读取配置
	config.Setup(configPath)
	// 2. 初始化数据库链接
	database.Setup()
	//test mongo
	//err := orm.EloquentMongo.Database()
	//if err!=nil{
	//	logger.Warn(err)
	//}
	// 3. 启动异步任务队列
	//go task.Start()
	//go cronjob.PingAgentService()
	//go task.StartWorker()

}

func run() error {
	if mode != "" {
		config.Mode = mode
	}
	if config.Mode == string(config.Prod) {
		gin.SetMode(gin.ReleaseMode)
	}

	r := router.InitRouter()

	defer func() {
		db, err := database.MysqlDB.DB()
		if err != nil {
			klog.Info(err)
		}
		err = db.Close()
		if err != nil {
			klog.Info(err)
		}
	}()

	//var r http.Handler
	srv := &http.Server{
		Addr:    config.ApplicationConfig.Host + ":" + config.ApplicationConfig.Port,
		Handler: r,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			klog.Fatalf("listen: %s\n", err.Error())
		}
	}()
	klog.Infof("%s Server Run http://%s:%s/ \r\n",
		tools.GetCurrentTimeStr(),
		config.ApplicationConfig.Host,
		config.ApplicationConfig.Port)
	klog.Infof("%s Swagger URL http://%s:%s/swagger/index.html \r\n",
		tools.GetCurrentTimeStr(),
		config.ApplicationConfig.Host,
		config.ApplicationConfig.Port)
	klog.Infof("%s Enter Control + C Shutdown Server \r\n", tools.GetCurrentTimeStr())
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	klog.Infof("%s Shutdown Server ... \r\n", tools.GetCurrentTimeStr())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		klog.Fatal("Server Shutdown:", err)
	}
	klog.Info("Server exiting")
	return nil
}
