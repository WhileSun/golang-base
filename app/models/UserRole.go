package models

import "github.com/whilesun/go-admin/app/po"

type RUserRole struct {
	po.RUserRole
}

func NewUserRole() *RUserRole {
	return &RUserRole{}
}

func (m *RUserRole) GetUserRoleIds(userId int) []string{
	var roleIds []string
	db.Model(&RUserRole{}).Select("role_id").Where("user_id=?",userId).Scan(&roleIds)
	return roleIds
}