package gcasbin

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/whilesun/go-admin/pkg/core/gconfig"
	"github.com/whilesun/go-admin/pkg/core/gdb"
	"log"
	"path/filepath"
	"runtime"
)

type CasbinConfig struct {
	Prefix    string
	TableName string
	AdminRole string
	ConfPath  string
}

func New(casbinKey string) *casbin.Enforcer {
	if casbinKey == "" {
		casbinKey = "casbin"
	}
	casbinConfig := &CasbinConfig{
		Prefix:    "r",
		TableName: "role_policy",
		AdminRole: "super_admin",
		ConfPath:  "rbac_model.conf",
	}
	gconfig.Config.UnmarshalKey(casbinKey, casbinConfig)
	return casbinConfig.initCasbin()
}

func (casbinConfig *CasbinConfig) initCasbin() *casbin.Enforcer {
	db := gdb.New("")
	adapter, err := gormadapter.NewAdapterByDBUseTableName(db, casbinConfig.Prefix, casbinConfig.TableName)
	if err != nil {
		log.Fatalf("casbin gormadapter error: %s", err.Error())
	}
	_, file, _, _ := runtime.Caller(1)
	enforcer, err := casbin.NewEnforcer(filepath.Dir(file)+"/"+casbinConfig.ConfPath, adapter)
	if err != nil {
		log.Fatalf("casbin NewEnforcer error: %s", err.Error())
	}
	enforcer.LoadPolicy()
	enforcer.AddFunction("checkSuperUser", func(args ...interface{}) (interface{}, error) {
		// 获取用户名
		username := args[0].(string)
		// 检查用户名的角色是否为super_admin
		ok, _ := enforcer.HasRoleForUser(username, casbinConfig.AdminRole)
		return ok, nil
	})
	return enforcer
}
