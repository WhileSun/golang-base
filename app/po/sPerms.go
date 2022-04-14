package po

//SPerms 节点表
type SPerms struct {
	BaseField
	Name      string `json:"name"`
	PagePerms string `json:"page_perms"`
	DataPerms string `json:"data_perms"`
}
