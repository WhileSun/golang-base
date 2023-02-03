package api

import "github.com/whilesun/go-admin/app/service"

type SysApiGroup struct {
	SysGlobalApi SysGlobalApi
	SysMenuApi   SysMenuApi
	SysPermsApi  SysPermsApi
	SysRoleApi   SysRoleApi
	SysUserApi   SysUserApi
}

type ApiGroup struct {
	SysApiGroup SysApiGroup
}

var ApiGroupApp = new(ApiGroup)

// service type
var (
	sysUserService  = service.ServiceGroupApp.SysServiceGroup.SysUserService
	sysRoleService  = service.ServiceGroupApp.SysServiceGroup.SysRoleService
	sysMenuService  = service.ServiceGroupApp.SysServiceGroup.SysMenuService
	sysPermsService = service.ServiceGroupApp.SysServiceGroup.SysPermsService
)
