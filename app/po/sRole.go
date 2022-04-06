package po

//SRole 角色表
type SRole struct {
	BaseField
	RoleName     string `json:"role_name"`
	RoleIdentity string `json:"role_identity"`
	Sort         int    `json:"sort"`
	Status       bool   `json:"status"`
	PermsIds     string `json:"perms_ids"`
}