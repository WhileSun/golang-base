package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/whilesun/go-admin/app/service"
	"github.com/whilesun/go-admin/pkg/e"
	"github.com/whilesun/go-admin/pkg/utils/gconvert"
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
		userInfo, ok := service.NewUserAuth().VerifyLogin(token)
		if !ok {
			e.New(con).Msg(e.ERROR_LOGIN_AUTH)
			con.Abort()
			return
		}
		con.Set("userId", gconvert.StrToInt(userInfo["user_id"]))
		con.Set("username", userInfo["username"])
		con.Set("userToken", token)
	}
}
