package dto

type AddRole struct {
	RoleName     string `form:"role_name"  binding:"required" label:"角色名称"`
	Sort         *int   `form:"sort" binding:"required,gt=0" label:"角色排序"`
	Status       *bool  `form:"status" binding:"required" label:"角色状态"`
	PermsIds     string `form:"perms_ids" binding:"" label:"菜单权限"`
}

type UpdateRole struct {
	Id int `form:"id"  binding:"required" label:"ID"`
	RoleName     string `form:"role_name"  binding:"required" label:"角色名称"`
	Sort         *int   `form:"sort" binding:"required,gt=0" label:"角色排序"`
	Status       *bool  `form:"status" binding:"required" label:"角色状态"`
	PermsIds     string `form:"perms_ids" binding:"" label:"菜单权限"`
}
