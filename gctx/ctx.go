package gctx

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/whilesun/go-admin/pkg/core/gconfig"
	"github.com/whilesun/go-admin/pkg/core/gdb"
	"github.com/whilesun/go-admin/pkg/core/glog"
	"github.com/whilesun/go-admin/pkg/core/gredis"
	"github.com/whilesun/go-admin/pkg/helper/gcaptcha"
	"github.com/whilesun/go-admin/pkg/helper/gjwt"
	"gorm.io/gorm"
)

var (
	Logger   *logrus.Logger
	Db       *gorm.DB
	GRedis   *gredis.GRedis
	GCaptcha *gcaptcha.CaptchaConfig
	GConfig  *viper.Viper
	GJwt     *gjwt.JwtConfig
)

func init() {
	fmt.Println("\r\nCTX-----------start")
	Db = gdb.New("")
	Logger = glog.New("log")
	GRedis = gredis.New("redis")
	GCaptcha = gcaptcha.New("captcha")
	GJwt = gjwt.New("jwt")
	GConfig = gconfig.Config
	fmt.Println("CTX-----------end")
}
