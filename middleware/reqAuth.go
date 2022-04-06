package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/whilesun/go-admin/app/service"
	"github.com/whilesun/go-admin/pkg/e"
	"strings"
)

func ReqAuth() gin.HandlerFunc {
	return func(con *gin.Context) {
		userAuthService := service.NewUserAuth()
		isSuper, userRoles := userAuthService.GetRoles(con.GetInt("userId"))
		if isSuper == 0 {
			path := con.Request.RequestURI
			path = strings.Replace(path, "/api/", "", 1)
			ok := userAuthService.ExistsPerm(path, userRoles)
			if !ok {
				e.New(con).Msg(e.ERROR_API_PERMS)
				con.Abort()
				return
			}
		}
	}
}
