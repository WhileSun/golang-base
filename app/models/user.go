package models

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/whilesun/go-admin/app/po"
	"github.com/whilesun/go-admin/app/vo"
	"github.com/whilesun/go-admin/pkg/utils/grequest"
)

// SUser 用户表
type SUser struct {
	po.SUser
}

func NewUser() *SUser {
	return &SUser{}
}

//CheckExist 监测用户是否存在
func (m *SUser) CheckExist(username string, password string) *SUser {
	db.Select("id,status,username,realname").Where("username=? and password=?", username, password).Find(m)
	return m
}

func (m *SUser) GetInfo(id int) *SUser {
	db.Select("id,username,realname").Where("id = ?", id).Find(m)
	return m
}

func (m *SUser) GetList(req *gin.Context) []*vo.UserList {
	bindParams := make(map[string]interface{}, 0)
	limit := grequest.PageLimit(req, bindParams)
	searchParams := map[string]string{"u*realname": "like", "u*username": "like", "u*status": "bool"}
	where := grequest.ParamsWhere(req, searchParams, bindParams)
	users := make([]*vo.UserList,0)
	db.Raw(fmt.Sprintf(`SELECT
			u.id,
			u.username,
			u.realname,
			u.status,
			u.created_at,
			u.updated_at,
			js.role_ids_str,
			js,role_names 
	FROM 
		s_user u
		left join
		(select 
					u.id,string_agg(r.id::text,',') as role_ids_str,string_agg(r.role_name,',') as role_names
		from s_user u
					left join r_user_role ur on u.id = ur.user_id
					left join s_role r on ur.role_id= r.id
					GROUP BY 1
		) js on u.id = js.id
		where 1=1 %s
		order by u.id desc %s`, where, limit), bindParams).Scan(&users)
	return users
}

//CheckUsernameExist 判断用户账号是否已经创建
func (m *SUser) CheckUsernameExist() int {
	var id int
	db.Model(&SUser{}).Select("id").Where("username = ?", m.Username).Scan(&id)
	return id
}

func (m *SUser) CheckRealnameExist() int {
	var id int
	db.Model(&SUser{}).Select("id").Where("realname = ?", m.Realname).Scan(&id)
	return id
}

//GetRoles 获取用户有那些管理员权限
func (m *SUser) GetRoles(userId int) []string {
	roles := make([]string, 0)
	db.Raw(`SELECT role_identity FROM 
		s_user u 
	left join r_user_role ur on u.id = ur.user_id
	left join s_role r on ur.role_id = r.id
	where u.id = ?
	`, userId).Scan(&roles)
	return roles
}
