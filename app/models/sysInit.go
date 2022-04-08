package models

import (
	"github.com/whilesun/go-admin/app/po"
	"github.com/whilesun/go-admin/pkg/gconf"
	"github.com/whilesun/go-admin/pkg/gcrypto"
	"github.com/whilesun/go-admin/pkg/utils/gtools"
	"gorm.io/gorm"
	"log"
)

type SysInit struct {
}

func NewSysInit() *SysInit {
	return &SysInit{}
}

func (m *SysInit) Run() {
	err := db.Transaction(func(tx *gorm.DB) error {
		var dbInit string
		tx.Model(&SSetting{}).Select("value").Where("types=?", "sys").Where("label = ?", "dbInit").Scan(&dbInit)
		if dbInit == "ok" {
			return nil
		}
		roleModel := NewRole()
		roleModel.RoleName = "超级管理员"
		superRoleIdentity := gtools.StringDefault(gconf.Config.GetString("app.roleSuperName"), "super_admin")
		roleModel.RoleIdentity = superRoleIdentity
		roleModel.Sort = 1
		roleModel.Status = true
		if err := tx.Create(&roleModel).Error; err != nil {
			return err
		}
		userModel := NewUser()
		superUserName := gtools.StringDefault(gconf.Config.GetString("app.userSuperName"), "system")
		userModel.Username = superUserName
		initPwd := gconf.Config.GetString("app.initPwd")
		userModel.Password = gcrypto.PwdEncode(gtools.StringDefault(initPwd, "a123456!"))
		userModel.Realname = "超级用户"
		userModel.Status = true
		if err := tx.Create(&userModel).Error; err != nil {
			return err
		}
		userRoleModel := NewUserRole()
		userRoleModel.UserId = userModel.Id
		userRoleModel.RoleId = roleModel.Id
		if err := tx.Create(&userRoleModel).Error; err != nil {
			return err
		}
		//创建初始化目录
		if err := initMenu(tx); err != nil {
			return err
		}
		settingModel := NewSetting()
		settingModel.Types = "sys"
		settingModel.Label = "dbInit"
		settingModel.Value = "ok"
		if err := tx.Create(&settingModel).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Fatalf("初始化数据库表失败,Error %s", err.Error())
	}
}

func initMenu(tx *gorm.DB) error {
	systemModel := &SMenu{po.SMenu{MenuName: "系统管理", Url: "/system", Icon: "icon-xitong", ParentId: 0, Sort: 100, MenuType: 1, Status: true, Show: true, IsSys: true}}
	if err := tx.Create(&systemModel).Error; err != nil {
		return err
	}
	userModel := &SMenu{po.SMenu{MenuName: "用户管理", Url: "/system/user", Icon: "icon-yonghuguanli_huaban", ParentId: systemModel.Id, Sort: 1, MenuType: 1, Status: true, Show: true, IsSys: true}}
	if err := tx.Create(&userModel).Error; err != nil {
		return err
	}
	userPermsModel := []*SMenu{
		{po.SMenu{MenuName: "列表", DataPerms: "user/list/get", PagePerms: "LIST", ParentId: userModel.Id, Sort: 1, MenuType: 2, Status: true, Show: true, IsSys: true}},
		{po.SMenu{MenuName: "添加", DataPerms: "user/add", PagePerms: "ADD", ParentId: userModel.Id, Sort: 2, MenuType: 2, Status: true, Show: true, IsSys: true}},
		{po.SMenu{MenuName: "修改", DataPerms: "user/update", PagePerms: "UPDATE", ParentId: userModel.Id, Sort: 3, MenuType: 2, Status: true, Show: true, IsSys: true}},
	}
	if err := tx.Create(&userPermsModel).Error; err != nil {
		return err
	}

	roleModel := &SMenu{po.SMenu{MenuName: "角色管理", Url: "/system/role", Icon: "icon-jiaoseshezhi", ParentId: systemModel.Id, Sort: 2, MenuType: 1, Status: true, Show: true, IsSys: true}}
	if err := tx.Create(&roleModel).Error; err != nil {
		return err
	}
	rolePermsModel := []*SMenu{
		{po.SMenu{MenuName: "列表", DataPerms: "role/list/get", PagePerms: "LIST", ParentId: roleModel.Id, Sort: 1, MenuType: 2, Status: true, Show: true, IsSys: true}},
		{po.SMenu{MenuName: "添加", DataPerms: "role/add", PagePerms: "ADD", ParentId: roleModel.Id, Sort: 2, MenuType: 2, Status: true, Show: true, IsSys: true}},
		{po.SMenu{MenuName: "修改", DataPerms: "role/update", PagePerms: "UPDATE", ParentId: roleModel.Id, Sort: 3, MenuType: 2, Status: true, Show: true, IsSys: true}},
	}
	if err := tx.Create(&rolePermsModel).Error; err != nil {
		return err
	}

	menuModel := &SMenu{po.SMenu{MenuName: "菜单管理", Url: "/system/menu", Icon: "icon-caidan", ParentId: systemModel.Id, Sort: 3, MenuType: 1, Status: true, Show: true, IsSys: true}}
	if err := tx.Create(&menuModel).Error; err != nil {
		return err
	}
	menuPermsModel := []*SMenu{
		{po.SMenu{MenuName: "列表", DataPerms: "menu/list/get", PagePerms: "LIST", ParentId: menuModel.Id, Sort: 1, MenuType: 2, Status: true, Show: true, IsSys: true}},
		{po.SMenu{MenuName: "添加", DataPerms: "menu/add", PagePerms: "ADD", ParentId: menuModel.Id, Sort: 2, Status: true, Show: true,MenuType: 2, IsSys: true}},
		{po.SMenu{MenuName: "修改", DataPerms: "menu/update", PagePerms: "UPDATE", ParentId: menuModel.Id, Sort: 3, MenuType: 2, Status: true, Show: true, IsSys: true}},
		{po.SMenu{MenuName: "删除", DataPerms: "menu/delete", PagePerms: "DELETE", ParentId: menuModel.Id, Sort: 4, MenuType: 2, Status: true, Show: true, IsSys: true}},
	}
	if err := tx.Create(&menuPermsModel).Error; err != nil {
		return err
	}
	return nil
}
