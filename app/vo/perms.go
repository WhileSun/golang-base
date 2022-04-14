package vo

type PermsFieldList struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	PagePerms string `json:"page_perms"`
}

//PermsMenuList 用菜单栏批量添加子节点
type PermsMenuList struct {
	Name      string `json:"name"`
	DataPerms string `json:"data_perms"`
	PagePerms string `json:"page_perms"`
}
