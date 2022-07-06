package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/whilesun/go-admin/app/service"
	"github.com/whilesun/go-admin/pkg/e"
	"github.com/whilesun/go-admin/pkg/utils/gtools"
	"strings"
)

func ReqAuth() gin.HandlerFunc {
	return func(con *gin.Context) {
		path := con.Request.RequestURI
		path = strings.Replace(path, "/api/", "", 1)
		ok := service.NewUserAuth().CheckRole(gtools.GetUserRoleIdents(con),path)
		if !ok {
			e.New(con).Msg(e.ERROR_API_PERMS)
			con.Abort()
			return
		}
	}
}
