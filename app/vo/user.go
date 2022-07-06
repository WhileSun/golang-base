package vo

import (
	"github.com/whilesun/go-admin/app/po"
)

type UserList struct {
	po.BaseField
	Username   string `json:"username"`
	Realname   string `json:"realname"`
	Status     bool   `json:"status"`
	RoleIdsStr string `json:"role_ids_str"`
	RoleNames  string `json:"role_names"`
}

type UserSession struct {
	po.BaseField
	Username   string `json:"username"`
	Realname   string `json:"realname"`
	RoleIdents string `json:"role_idents"`
}
