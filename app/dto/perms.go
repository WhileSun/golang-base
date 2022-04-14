package dto

type AddPerms struct {
	Name      string `form:"name"  binding:"required" label:"节点名称"`
	PagePerms string `form:"page_perms"  binding:"required" label:"操作权限标识"`
	DataPerms string `form:"data_perms" binding:"required" label:"数据权限标识"`
}

type UpdatePerms struct {
	Id int `form:"id"  binding:"required" label:"ID"`
	AddPerms
}

type DeletePerms struct {
	Id int `form:"id"  binding:"required" label:"ID"`
}
