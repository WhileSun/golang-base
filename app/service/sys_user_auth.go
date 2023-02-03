package service

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/whilesun/go-admin/gctx"
	"strings"
	"time"
)

type SysUserAuthService struct {
}

var SysUserAuthServiceApp = new(SysUserAuthService)

func (s *SysUserAuthService) CreateJwtToken(userId int) (token string, err error) {
	token, err = gctx.GJwt.CreateToken(jwt.MapClaims{"userId": userId})
	return
}

func (s *SysUserAuthService) ParseJwtToken(token string) (jwt.MapClaims, string, bool) {
	var newToken string
	userInfo, err := gctx.GJwt.ParseToken(token)
	//解析失败查找是否过期
	if err != nil {
		if strings.Contains(err.Error(), "expired") {
			nowTime := time.Now().Unix()
			if int64(userInfo["lastexp"].(float64)) >= nowTime {
				newToken, err = gctx.GJwt.CreateToken(userInfo)
			}
		}
	}
	if err != nil {
		return userInfo, newToken, false
	} else {
		return userInfo, newToken, true
	}
}
