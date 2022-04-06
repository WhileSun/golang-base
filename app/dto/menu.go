package dto

type AddMenu struct {
	MenuName  string `form:"menu_name"  binding:"required" label:"菜单名称"`
	MenuType  int16  `form:"menu_type"  binding:"required,gt=0,lte=2" label:"菜单类型"`
	Url       string `form:"url" label:"菜单链接 "`
	Icon      string `form:"icon"  label:"菜单图标"`
	ParentId  *int   `form:"parent_id" binding:"required,gte=0" label:"上级菜单"`
	Sort      int    `form:"sort" binding:"required,gt=0" label:"菜单排序"`
	Status    *bool  `form:"status" binding:"required" label:"菜单状态"`
	Show      *bool  `form:"show" binding:"required" label:"是否显示"`
	DataPerms string `form:"data_perms" label:"数据权限"`
	PagePerms string `form:"page_perms" label:"操作权限"`
}

type UpdateMenu struct {
	Id int `form:"id"  binding:"required" label:"ID"`
	AddMenu
}

type DeleteMenu struct {
	Id int `form:"id"  binding:"required" label:"ID"`
}
