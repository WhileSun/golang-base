package po

//SRole 角色表
type SRole struct {
	BaseField
	RoleName  string `json:"role_name"`
	RoleIdent string `json:"role_ident"`
	Sort      int    `json:"sort"`
	Status    bool   `json:"status"`
	PermsIds  string `json:"perms_ids"`
}
