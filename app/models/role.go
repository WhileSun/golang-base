package models

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/whilesun/go-admin/app/po"
	"github.com/whilesun/go-admin/app/vo"
	"github.com/whilesun/go-admin/pkg/utils/grequest"
	"strings"
)

type SRole struct {
	po.SRole
}

func NewRole() *SRole {
	return &SRole{}
}

func (m *SRole) GetRow(id int) *SRole {
	db.Model(&SRole{}).Where("id = ?", id).First(m)
	return m
}

func (m *SRole) Add() error {
	result := db.Create(m)
	return result.Error
}

func (m *SRole) Update() error {
	result := db.Model(m).Select("role_name", "sort", "status", "perms_ids").Updates(m)
	return result.Error
}

func (m *SRole) GetList(req *gin.Context) ([]*SRole, error) {
	bindParams := make(map[string]interface{}, 0)
	limit := grequest.PageLimit(req, bindParams)
	searchParams := map[string]string{"role_name": "like", "role_identity":"like","status": "bool"}
	where := grequest.ParamsWhere(req, searchParams, bindParams)
	roles := make([]*SRole, 0)
	db.Raw(fmt.Sprintf("SELECT * FROM s_role where 1=1 %s order by sort asc,id asc %s",
		where, limit), bindParams).Scan(&roles)
	return roles, nil
}

func (m *SRole) CheckRoleIdentityExist() int {
	var id int
	db.Model(&SRole{}).Select("id").Where("role_identity = ?", m.RoleIdentity).Scan(&id)
	return id
}

func (m *SRole) CheckRoleNameExist() int {
	var id int
	db.Model(&SRole{}).Select("id").Where("role_name = ?", m.RoleName).Scan(&id)
	return id
}

//GetRolePerms 获取某个权限标识下的角色权限
func (m *SRole) GetRolePerms(roleIdentity string) []string{
	var permsIds string
	db.Model(&SRole{}).Raw("SELECT perms_ids FROM s_role where role_identity =? and status = true",roleIdentity).Scan(&permsIds)
	perms := make([]string,0)
	if permsIds !=""{
		ids := strings.Split(permsIds,",")
		db.Model(&SMenu{}).Select("distinct data_perms").Where("menu_type=?",2).
			Where("status=?",true).
			Where("id in ?",ids).Scan(&perms)
	}
	return perms
}

//GetFieldList 获取主要字段列表
func (m *SRole) GetFieldList() []*vo.RoleFieldList{
	roles := make([] *vo.RoleFieldList,0)
	db.Model(&SRole{}).Raw(`SELECT id,role_name,role_identity FROM s_role order by sort asc,id asc`).Scan(&roles)
	return roles
}