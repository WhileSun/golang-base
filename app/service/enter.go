package service

type SysServiceGroup struct {
	SysUserService  SysUserService
	SysRoleService  SysRoleService
	SysMenuService  SysMenuService
	SysPermsService SysPermsService
}

type ServiceGroup struct {
	SysServiceGroup SysServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
