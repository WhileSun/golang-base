package service

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	uuid "github.com/satori/go.uuid"
	"github.com/whilesun/go-admin/app/models"
	"github.com/whilesun/go-admin/pkg/gconf"
	"github.com/whilesun/go-admin/pkg/gcrypto"
	"github.com/whilesun/go-admin/pkg/gsys"
)

type UserAuthService struct {
}

func NewUserAuth() *UserAuthService {
	return &UserAuthService{}
}

type loginSettingObj struct {
	TokenKey   string
	SessionKey string
	SessionEx  int
}

var loginSetting *loginSettingObj

func init() {
	loginSetting = &loginSettingObj{
		TokenKey:   "token",
		SessionKey: "user_session",
		SessionEx:  3600,
	}
	gconf.Config.UnmarshalKey("loginSetting", loginSetting)
}

func (s *UserAuthService) CreateLoginToken(userId int) string {
	return gcrypto.Md5Encode(fmt.Sprintf("%s_%d_%s", loginSetting.TokenKey, userId, uuid.NewV1()))
}

//SetSession 设置用户信息
func (s *UserAuthService) SetSession(user *models.SUser, token string) {
	redisConn := gsys.Redis.Get()
	defer redisConn.Close()
	key := fmt.Sprintf("%s_%s", loginSetting.SessionKey, token)
	redisConn.Do("HMSET", key, "user_id", user.Id, "username", user.Username,
		"realname", user.Realname)
	redisConn.Do("expire", key, loginSetting.SessionEx)
	//存储用户的session key
	tokenKey := fmt.Sprintf("%s_%d_%s", loginSetting.SessionKey, user.Id, token)
	redisConn.Do("SET", tokenKey, key, "EX", loginSetting.SessionEx)
}

//DelSession 删除用户信息
func (s *UserAuthService) DelSession(userId int, token string) {
	redisConn := gsys.Redis.Get()
	defer redisConn.Close()
	key := fmt.Sprintf("%s_%s", loginSetting.SessionKey, token)
	redisConn.Do("DEL", key)
	tokenKey := fmt.Sprintf("%s_%d_%s", loginSetting.SessionKey, userId, token)
	redisConn.Do("DEL", tokenKey)
}

//VerifyLogin 验证用户是否登录
func (s *UserAuthService) VerifyLogin(token string) (map[string]string, bool) {
	redisConn := gsys.Redis.Get()
	defer redisConn.Close()
	key := fmt.Sprintf("%s_%s", loginSetting.SessionKey, token)
	resp, _ := redis.StringMap(redisConn.Do("HGETALL", key))
	if len(resp) > 0 {
		return resp, true
	} else {
		return resp, false
	}
}

//GetRoles 获取用户所有的角色
func (s *UserAuthService) GetRoles(userId int) (int, []string) {
	redisConn := gsys.Redis.Get()
	defer redisConn.Close()
	roles := make([]string, 0)
	key := fmt.Sprintf("user_roles_%d", userId)
	exists, _ := redis.Int(redisConn.Do("EXISTS", key))
	if exists == 0 {
		roles = models.NewUser().GetRoles(userId)
		args := []interface{}{key}
		for _, role := range roles {
			args = append(args,role)
		}
		redisConn.Do("SADD", args...)
	} else {
		roles, _ = redis.Strings(redisConn.Do("SMEMBERS", key))
	}
	member := gconf.Config.GetString("app.roleSuperName")
	isSuper, _ := redis.Int(redisConn.Do("SISMEMBER", key, member))
	return isSuper, roles
}

//SetRolePerms 设置权限
func (s *UserAuthService) SetRolePerms(roleIdentity string){
	redisConn := gsys.Redis.Get()
	defer redisConn.Close()
	key := fmt.Sprintf("role_perms_%s", roleIdentity)
	exists, _ := redis.Int(redisConn.Do("EXISTS", key))
	if exists == 0{
		perms := models.NewRole().GetRolePerms(roleIdentity)
		args := []interface{}{key}
		if len(perms) == 0{
			args = append(args,"")
		}else{
			for _, perm := range perms {
				args = append(args,perm)
			}
		}
		redisConn.Do("SADD", args...)
	}
}

//ExistsPerm 鉴别用户是否存在权限
func (s *UserAuthService) ExistsPerm(perm string, roles []string) bool {
	redisConn := gsys.Redis.Get()
	defer redisConn.Close()
	userAuthService := NewUserAuth()
	for _, role := range roles {
		userAuthService.SetRolePerms(role)
		having, _ := redis.Int(redisConn.Do("SISMEMBER", fmt.Sprintf("role_perms_%s", role), perm))
		if having == 1 {
			return true
		}
	}
	return false
}

func (s *UserAuthService) DelRole(roleIdentity string){
	redisConn := gsys.Redis.Get()
	defer redisConn.Close()
	key:=fmt.Sprintf("role_perms_%s", roleIdentity)
	redisConn.Do("DEL", key)
}