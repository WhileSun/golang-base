package service

import (
	"errors"
	"fmt"
	"github.com/whilesun/go-admin/app/dto"
	"github.com/whilesun/go-admin/app/models"
	"github.com/whilesun/go-admin/pkg/gconf"
	"github.com/whilesun/go-admin/pkg/gsys"
	"github.com/whilesun/go-admin/pkg/utils/gconvert"
	"github.com/whilesun/go-admin/pkg/utils/gtools"
)

type RoleService struct {
}

func NewRole() *RoleService {
	return &RoleService{}
}

func checkRoleIdentityExist(roleModel *models.SRole) error {
	if id := roleModel.CheckRoleIdentityExist(); id > 0 {
		return errors.New(fmt.Sprintf("角色权限标识[%s]已经存在，请更换！", roleModel.RoleIdentity))
	}
	return nil
}

func checkRoleNameExist(roleModel *models.SRole) error {
	if id := roleModel.CheckRoleNameExist(); id > 0 {
		return errors.New(fmt.Sprintf("角色名称[%s]已经存在，请更换！", roleModel.RoleName))
	}
	return nil
}

func (s *RoleService) Add(params dto.AddRole) error {
	roleModel := models.NewRole()
	gconvert.StructCopy(params, roleModel)
	if err := checkRoleIdentityExist(roleModel); err != nil {
		return err
	}
	if err := checkRoleNameExist(roleModel); err != nil {
		return err
	}
	err := roleModel.Add()
	if err != nil {
		gsys.Logger.Error("添加角色失败—>", err.Error())
		return errors.New("添加角色失败！")
	}
	return nil
}

func (s *RoleService) Update(params dto.UpdateRole) error {
	odlRoleModel := models.NewRole()
	odlRoleModel.GetRow(params.Id)
	if odlRoleModel.Id == 0 {
		return errors.New("需要更新的角色不存在，请确认！")
	}
	roleSuperName := gtools.StringDefault(gconf.Config.GetString("app.roleSuperName"), "super_admin")
	if odlRoleModel.RoleIdentity == roleSuperName {
		return errors.New("超级管理员角色不支持编辑！")
	}
	//更新角色
	roleModel := models.NewRole()
	gconvert.StructCopy(params, roleModel)
	//if odlRoleModel.RoleIdentity != roleModel.RoleIdentity {
	//	if err := checkRoleIdentityExist(roleModel); err != nil {
	//		return err
	//	}
	//}
	if roleModel.RoleName != roleModel.RoleName {
		if err := checkRoleNameExist(roleModel); err != nil {
			return err
		}
	}
	err := roleModel.Update()
	if err != nil {
		gsys.Logger.Error("编辑角色失败—>", err.Error())
		return errors.New("更新角色失败！")
	}
	NewUserAuth().DelRolePerms(odlRoleModel.RoleIdentity)
	return nil
}
