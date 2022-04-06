package grequest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/whilesun/go-admin/pkg/utils/gconvert"
	"strings"
)

//PageLimit 分页
func PageLimit(req *gin.Context, bindParams map[string]interface{}) string {
	page:= gconvert.StrToInt(req.PostForm("page"))
	if page == 0 {
		page = 1
	}
	pageSize := gconvert.StrToInt(req.PostForm("page_size"))
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	offset := (page - 1) * pageSize
	bindParams["offset"] = offset
	bindParams["limit"] = pageSize
	return "offset @offset limit @limit"
}

//ParamsWhere where 条件
func ParamsWhere(req *gin.Context, params map[string]string, bindParams map[string]interface{}) string {
	sql := ""
	for param, paramType := range params {
		value := req.PostForm(param)
		value = strings.Trim(value, " ")
		if value == "" {
			continue
		}
		param = strings.Replace(param, "*", ".", 1)
		var newValue interface{}
		switch paramType {
		case "like":
			newValue = "%" + value + "%"
			sql += fmt.Sprintf(" and %s like @%s", param, param)
		case "string":
			newValue = value
			sql += fmt.Sprintf(" and %s=@%s", param, param)
		case "bool":
			newValue = gconvert.StrToBool(value)
			sql += fmt.Sprintf(" and %s=@%s", param, param)
		case "int":
			newValue = gconvert.StrToInt(value)
			sql += fmt.Sprintf(" and %s=@%s", param, param)
		default:
			newValue = value
			sql += fmt.Sprintf(" and %s=@%s", param, param)
		}
		bindParams[param] = newValue
	}
	return sql
}