package vo

// SysRoleModel s_role 角色表
type SysRoleModel struct {
	BaseField
	RoleName  string `json:"role_name"`
	RoleIdent string `json:"role_ident"`
	Sort      int    `json:"sort"`
	Status    bool   `json:"status"`
	PermsIds  string `json:"perms_ids"`
}

type RoleFieldList struct {
	Id        int    `json:"id"`
	RoleName  string `json:"role_name"`
	RoleIdent string `json:"role_ident"`
}
