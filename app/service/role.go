package service

import (
	"errors"
	"fmt"
	"github.com/whilesun/go-admin/app/dto"
	"github.com/whilesun/go-admin/app/models"
	"github.com/whilesun/go-admin/pkg/gconf"
	"github.com/whilesun/go-admin/pkg/gcrypto"
	"github.com/whilesun/go-admin/pkg/gsys"
	"github.com/whilesun/go-admin/pkg/utils/gconvert"
	"github.com/whilesun/go-admin/pkg/utils/gtools"
	"math/rand"
	"time"
)

type Role struct {
}

func NewRole() *Role {
	return &Role{}
}

//checkRoleNameExist 检测角色名是否存在
func (s *Role) checkRoleNameExist(roleName string) error {
	if id := models.NewRole().CheckRoleNameExist(roleName); id > 0 {
		return errors.New(fmt.Sprintf("角色名称[%s]已经存在，请更换！", roleName))
	}
	return nil
}

func (s *Role) Add(params dto.AddRole) error {
	roleModel := models.NewRole()
	gconvert.StructCopy(params, roleModel)
	roleModel.RoleIdent = gcrypto.Md5Encode16(fmt.Sprintf("%d-%d",time.Now().UnixNano(),rand.Intn(99999)))
	if err := NewRole().checkRoleNameExist(roleModel.RoleName); err != nil {
		return err
	}
	err := roleModel.Add()
	if err != nil {
		gsys.Logger.Error("添加角色失败—>", err.Error())
		return errors.New("添加角色失败！")
	}
	return nil
}

func (s *Role) Update(params dto.UpdateRole) error {
	oldRoleModel := models.NewRole()
	oldRoleModel.GetInfo(params.Id)
	if oldRoleModel.Id == 0 {
		return errors.New("需要更新的角色不存在，请确认！")
	}
	roleSuperName := gtools.StringDefault(gconf.Config.GetString("app.roleSuperName"), "super_admin")
	if oldRoleModel.RoleIdent == roleSuperName {
		return errors.New("超级管理员角色不支持编辑！")
	}
	//更新角色
	roleModel := models.NewRole()
	gconvert.StructCopy(params, roleModel)
	if oldRoleModel.RoleName != roleModel.RoleName {
		if err := NewRole().checkRoleNameExist(roleModel.RoleName); err != nil {
			return err
		}
	}
	err := roleModel.Update()
	if err != nil {
		gsys.Logger.Error("编辑角色失败—>", err.Error())
		return errors.New("更新角色失败！")
	}
	//角色关闭或修改权限，清空缓存
	if roleModel.PermsIds != oldRoleModel.PermsIds || (roleModel.Status == false && roleModel.Status != oldRoleModel.Status){
		NewUserAuth().DelRolePerms(oldRoleModel.RoleIdent)
	}
	return nil
}
