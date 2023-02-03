package gdb

import (
	"github.com/whilesun/go-admin/pkg/core/gconfig"
	"gorm.io/gorm/logger"
	"log"
	"strings"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// DatabaseConfig 数据库连接设置
type DatabaseConfig struct {
	Type        string
	Host        string
	Port        string
	User        string
	Password    string
	Name        string
	TablePrefix string
	Charset     string
	MaxIdleConn int
	MaxOpenConn int
	Log         bool
}

var (
	dbURI     string
	dialector gorm.Dialector
)

// New 初始化数据库连接
func New(dbKey string) *gorm.DB {
	if dbKey == "" {
		dbKey = "database"
	}
	databaseConfig := &DatabaseConfig{
		Type:        "postgres",
		Host:        "127.0.0.1",
		Port:        "5432",
		User:        "root",
		Password:    "root",
		TablePrefix: "",
		Charset:     "utf8",

		MaxIdleConn: 20,
		MaxOpenConn: 200,
		Log:         false,
	}
	gconfig.Config.UnmarshalKey(dbKey, databaseConfig)
	return databaseConfig.run()
}

// run 初始化运行
func (databaseConfig *DatabaseConfig) run() *gorm.DB {
	dbType := strings.ToLower(databaseConfig.Type)
	if dbType == "mysql" {
		databaseConfig.mySqInit()
	} else if dbType == "postgres" {
		databaseConfig.postGresInit()
	} else {
		log.Fatalf("config database type [%s] is not setting", dbType)
	}
	var newLogger logger.Interface
	if databaseConfig.Log {
		newLogger = logger.New(
			NewWriter(),
			//log.New(os.Stdout, "\r\n", log.LstdFlags), // gorm io writer（日志输出的目标，前缀和日志包含的内容——译者注）
			logger.Config{
				SlowThreshold:             time.Second, // 慢 SQL 阈值
				LogLevel:                  logger.Info, // 日志级别
				IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
				Colorful:                  false,       // 禁用彩色打印
			},
		)
	}
	conn, err := gorm.Open(dialector, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{TablePrefix: databaseConfig.TablePrefix, SingularTable: true},
		Logger:         newLogger,
	})
	if err != nil {
		log.Fatalf("database gorm conn failed,Error: %s", err.Error())
	}
	sqlDB, err := conn.DB()
	if err != nil {
		log.Fatalf("database connect server failed,Error: %s", err.Error())
	}
	sqlDB.SetMaxIdleConns(databaseConfig.MaxIdleConn) // 空闲进程数
	sqlDB.SetMaxOpenConns(databaseConfig.MaxOpenConn) // 最大进程数
	sqlDB.SetConnMaxLifetime(time.Second * 600)       // 设置了连接可复用的最大时间
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("database ping is failed,Error:%s", err.Error())
	}
	return conn
}
