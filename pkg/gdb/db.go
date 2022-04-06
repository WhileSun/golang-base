package gdb

import (
	"fmt"
	"github.com/whilesun/go-admin/pkg/gconf"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"strings"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

//DatabaseConfigObj 数据库连接设置
type DatabaseConfigObj struct {
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

var databaseConfig *DatabaseConfigObj
var dbURI string
var dialector gorm.Dialector
var db *gorm.DB

func Run() {
	if db != nil{
		return
	}
	databaseConfig = &DatabaseConfigObj{
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
	gconf.Config.UnmarshalKey("database", databaseConfig)
	db = initDB()
}

func Get() *gorm.DB {
	return db
}

// initDB 初始化连接
func initDB() *gorm.DB {
	dbType := strings.ToLower(databaseConfig.Type)
	if dbType == "mysql" {
		mySqInit()
	} else if dbType == "postgres" {
		postGresInit()
	} else {
		log.Fatalf("config database type [%s] is not setting", dbType)
	}
	var newLogger logger.Interface
	if databaseConfig.Log {
		newLogger = logger.New(
			//glog.NewDbWriter(),
			log.New(os.Stdout, "\r\n", log.LstdFlags), // gorm io writer（日志输出的目标，前缀和日志包含的内容——译者注）
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
		fmt.Println("database ping is failed")
	}
	return conn
}
