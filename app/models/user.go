package models

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/whilesun/go-admin/app/po"
	"github.com/whilesun/go-admin/app/vo"
	"github.com/whilesun/go-admin/pkg/utils/grequest"
)

type SUser struct {
	po.SUser
}

func NewUser() *SUser {
	return &SUser{}
}

//GetLoginInfo 检测账号密码是否正确
func (m *SUser) GetLoginInfo(username string, password string) *SUser {
	db.Select("id,status").Where("username=? and password=?", username, password).Find(m)
	return m
}

//GetSessionInfo 获取用户session存储信息
func (m *SUser) GetSessionInfo(userId int) *vo.UserSession {
	userSession := &vo.UserSession{}
	db.Model(&SUser{}).Select("id,username,realname").Where("id=?",userId).Scan(userSession)
	db.Raw(`SELECT
			string_agg(sl.role_ident, ',') as role_idents
		FROM 
		r_user_role r
		left join s_role sl on r.role_id = sl.id
		where r.user_id = ?`, userId).Scan(&userSession.RoleIdents)
	return userSession
}

func (m *SUser) GetPasswd(userId int) string {
	var password string
	db.Model(&SUser{}).Select("password").Where("id=?", userId).Scan(&password)
	return password
}

func (m *SUser) UpdatePasswd() error {
	return db.Model(m).Select("password").Updates(m).Error
}

func (m *SUser) GetInfo(id int) *SUser {
	db.Select("id,username,realname,status").Where("id = ?", id).Find(m)
	return m
}

func (m *SUser) GetList(req *gin.Context) []*vo.UserList {
	bindParams := make(map[string]interface{}, 0)
	limit := grequest.PageLimit(req, bindParams)
	searchParams := map[string]string{"u*realname": "like", "u*username": "like", "u*status": "bool"}
	where := grequest.ParamsWhere(req, searchParams, bindParams)
	users := make([]*vo.UserList, 0)
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
func (m *SUser) CheckUsernameExist(username string) int {
	var id int
	db.Model(&SUser{}).Select("id").Where("username = ?",username).Scan(&id)
	return id
}

func (m *SUser) CheckRealnameExist(realname string) int {
	var id int
	db.Model(&SUser{}).Select("id").Where("realname = ?", realname).Scan(&id)
	return id
}
