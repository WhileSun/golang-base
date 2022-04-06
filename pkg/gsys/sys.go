package gsys

import (
	goredis "github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
	"github.com/whilesun/go-admin/pkg/gdb"
	"github.com/whilesun/go-admin/pkg/glog"
	"github.com/whilesun/go-admin/pkg/gredis"
	"gorm.io/gorm"
)

var (
	Logger   *logrus.Logger
	Db       *gorm.DB
	//Enforcer *casbin.Enforcer
	Redis    *goredis.Pool
)

func init(){
	gdb.Run()
	gredis.Run()
	glog.Run()
	//gcasbin.Run()

	Db = gdb.Get()
	Logger = glog.Get()
	Redis = gredis.GetPool()
	//Enforcer = gcasbin.Get()
}