package database

import (
	"autoOps/config"
	"bytes"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql" //加载mysql
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"k8s.io/klog/v2"
	"strconv"
)

var (
	DbType   string
	Host     string
	Port     int
	Name     string
	Username string
	Password string
)

func (e *Mysql) Setup() {

	var err error
	var db Database

	db = new(Mysql)
	MysqlConn = db.GetConnect()
	MysqlDB, err = db.Open(MysqlConn)
	if config.Prod != config.Mode {
		MysqlDB = MysqlDB.Debug()
	}
	if err != nil {
		klog.Fatalf("%s connect error %v", DbType, err)
	} else {
		klog.Infof("%s connect success!", DbType)
	}

	if MysqlDB.Error != nil {
		klog.Fatalf("database error %v", MysqlDB.Error)
	}

	// 是否开启详细日志记录
	//orm.Eloquent.LogMode(viper.GetBool("settings.gorm.logMode"))

	// 设置最大打开连接数
	s, err := MysqlDB.DB()
	if err != nil {
		klog.Fatalf("database error %v", err)
	}

	s.SetMaxOpenConns(viper.GetInt("settings.gorm.maxOpenConn"))
	// 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用
	s.SetMaxIdleConns(viper.GetInt("settings.gorm.maxIdleConn"))
}

type Mysql struct {
}

func (e *Mysql) Open(dsn string) (db *gorm.DB, err error) {
	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.LogLevel(viper.GetInt("settings.gorm.logMode"))),
		PrepareStmt:                              true,
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "ops_", // table name prefix, table for `User` would be `t_users`
			SingularTable: true,  // use singular table name, table for `User` would be `user` with this option enabled
			//NameReplacer:  strings.NewReplacer("Id", "ID"), // use name replacer to change struct/field name before convert it to db name
		},
	})
}

func (e *Mysql) GetConnect() string {

	DbType = config.DatabaseConfig.Dbtype
	Host = config.DatabaseConfig.Host
	Port = config.DatabaseConfig.Port
	Name = config.DatabaseConfig.Name
	Username = config.DatabaseConfig.Username
	Password = config.DatabaseConfig.Password

	var conn bytes.Buffer
	conn.WriteString(Username)
	conn.WriteString(":")
	conn.WriteString(Password)
	conn.WriteString("@tcp(")
	conn.WriteString(Host)
	conn.WriteString(":")
	conn.WriteString(strconv.Itoa(Port))
	conn.WriteString(")")
	conn.WriteString("/")
	conn.WriteString(Name)
	conn.WriteString("?charset=utf8&parseTime=True&loc=Local&timeout=10000ms")
	return conn.String()
}
