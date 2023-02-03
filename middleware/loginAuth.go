package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/whilesun/go-admin/app/service"
	"github.com/whilesun/go-admin/pkg/utils/e"
)

func LoginAuth() gin.HandlerFunc {
	return func(con *gin.Context) {
		token := ""
		authBearer := con.Request.Header.Get("Authorization")
		if authBearer != "" && len(authBearer) >= 7 {
			token = authBearer[7:]
		}
		if token == "" {
			e.New(con).Msg(e.ERROR_LOGIN_AUTH)
			con.Abort()
			return
		}
		userInfo, newToken, ok := service.SysUserAuthServiceApp.ParseJwtToken(token)
		if !ok {
			e.New(con).Msg(e.ERROR_LOGIN_AUTH)
			con.Abort()
			return
		}
		if newToken != "" {
			con.Header("newToken", newToken)
		}
		//con.Set("userSession", userInfo)
		con.Set("userId", int(userInfo["userId"].(float64)))
		//con.Set("username", userInfo["username"])
		//con.Set("userToken", sessionKey)
	}
}
