package vo

// SysUserModel s_user用户表
type SysUserModel struct {
	BaseField
	Username string `json:"username"`
	Password string `json:"password"`
	Realname string `json:"realname"`
	Status   bool   `json:"status"`
}

// SysUserRoleModel r_user_role 用户角色关联表
type SysUserRoleModel struct {
	UserId int `json:"user_id"`
	RoleId int `json:"role_id"`
}

type UserList struct {
	BaseField
	Username   string `json:"username"`
	Realname   string `json:"realname"`
	Status     bool   `json:"status"`
	RoleIdsStr string `json:"role_ids_str"`
	RoleNames  string `json:"role_names"`
}

type UserSession struct {
	BaseField
	Username   string `json:"username"`
	Realname   string `json:"realname"`
	RoleIdents string `json:"role_idents"`
}
