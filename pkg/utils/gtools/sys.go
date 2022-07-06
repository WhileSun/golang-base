package gtools

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func GetUserRoleIdents(req *gin.Context) []string{
	userSession := req.GetStringMapString("userSession")
	roleIdents := strings.Split(userSession["role_idents"],",")
	return roleIdents
}