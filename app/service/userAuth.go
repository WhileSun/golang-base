package service

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/whilesun/go-admin/app/models"
	gsys2 "github.com/whilesun/go-admin/gctx"
	gconf2 "github.com/whilesun/go-admin/pkg/core/gconfig"
	gcrypto2 "github.com/whilesun/go-admin/pkg/utils/gcrypto"
	"github.com/whilesun/go-admin/pkg/utils/gtools"
	"strconv"
)

type UserAuthService struct {
}

func NewUserAuth() *UserAuthService {
	return &UserAuthService{}
}

type loginSettingObj struct {
	AesKey      string
	AesVi       string
	TokenEx     int
	SessionName string
	SessionEx   int
}

var loginSetting *loginSettingObj

func init() {
	loginSetting = &loginSettingObj{
		AesKey:      "EfoBjp9cQtjaEVGiQUu8RsXqW5dRrLGS",
		AesVi:       "6bVS6Yym5lVPLmFW",
		TokenEx:     3600,
		SessionName: "user_session",
		SessionEx:   604800, //7天
	}
	gconf2.Config.UnmarshalKey("loginSetting", loginSetting)
}

func (s *UserAuthService) TokenEncode(userId int) (string, string) {
	sessionKey := strconv.Itoa(userId) + ":" + gcrypto2.Md5Encode(fmt.Sprintf("%d_%s", userId, uuid.NewV1()))
	return sessionKey, gcrypto2.AesEncode(sessionKey, loginSetting.AesKey, loginSetting.AesVi)
}

func (s *UserAuthService) TokenDecode(token string) string {
	return gcrypto2.AesDecode(token, loginSetting.AesKey, loginSetting.AesVi)
}

// SetSession 设置用户信息
func (s *UserAuthService) SetSession(userModel *models.SUser, sessionKey string) {
	key := fmt.Sprintf("%s:%s", loginSetting.SessionName, sessionKey)
	userSession := models.NewUser().GetSessionInfo(userModel.Id)
	gsys2.GRedis.Hmset(key, loginSetting.SessionEx, "user_id", userSession.Id, "username", userSession.Username,
		"realname", userSession.Realname, "role_idents", userSession.RoleIdents)
}

// DelAllSession 删除用户所有登录session
func (s *UserAuthService) DelAllSession(userId int) {
	pattern := fmt.Sprintf("%s:%d:%s", loginSetting.SessionName, userId, "*")
	sessionKeys, _ := gsys2.GRedis.Keys(pattern)
	for _, sessionKey := range sessionKeys {
		gsys2.GRedis.Del(sessionKey)
	}
}

// DelSession 删除用户信息
func (s *UserAuthService) DelSession(token string) {
	key := fmt.Sprintf("%s:%s", loginSetting.SessionName, token)
	gsys2.GRedis.Del(key)
}

// VerifyLogin 验证用户是否登录
func (s *UserAuthService) VerifyLogin(token string) (map[string]string, bool) {
	key := fmt.Sprintf("%s:%s", loginSetting.SessionName, token)
	resp, _ := gsys2.GRedis.Hgetall(key)
	if len(resp) > 0 {
		return resp, true
	} else {
		return resp, false
	}
}

func (s *UserAuthService) CheckIsSuper(roleIdents []string) bool {
	roleSuperName := gconf2.Config.GetString("app.roleSuperName")
	isSuper := gtools.InStrArray(roleSuperName, roleIdents)
	return isSuper
}

// CheckRole 验证用户是否有权限
func (s *UserAuthService) CheckRole(roleIdents []string, perms string) bool {
	userAuthService := NewUserAuth()
	isSuper := userAuthService.CheckIsSuper(roleIdents)
	//超级管理员
	if isSuper {
		return true
	}
	for _, roleIdent := range roleIdents {
		userAuthService.SetRolePerms(roleIdent)
		having, _ := gsys2.GRedis.Sismember(fmt.Sprintf("role_perms:%s", roleIdent), perms)
		if having == 1 {
			return true
		}
	}
	return false
}

// SetRolePerms 设置权限
func (s *UserAuthService) SetRolePerms(roleIdent string) {
	key := fmt.Sprintf("role_perms:%s", roleIdent)
	exists, _ := gsys2.GRedis.Exists(key)
	if exists == 0 {
		permsList := models.NewRole().GetRolePerms(roleIdent)
		gsys2.GRedis.Sadd(key, 0, permsList...)
	}
}

func (s *UserAuthService) DelRolePerms(roleIdent string) {
	key := fmt.Sprintf("role_perms:%s", roleIdent)
	gsys2.GRedis.Del(key)
}
