package service

import (
	"errors"
	"fmt"
	"github.com/whilesun/go-admin/app/dto"
	"github.com/whilesun/go-admin/app/models"
	"github.com/whilesun/go-admin/pkg/gsys"
	"github.com/whilesun/go-admin/pkg/utils/gconvert"
	"strings"
)

type Perms struct {
}

func NewPerms() *Perms {
	return &Perms{}
}

func (s *Perms) checkPagePermsExist(pagePerms string) error {
	if id := models.NewPerms().CheckPagePermsExist(pagePerms); id > 0 {
		return errors.New(fmt.Sprintf("节点操作权限标识[%s]已经存在，请更换！", pagePerms))
	}
	return nil
}

func (s *Perms) checkNameExist(name string) error {
	if id := models.NewPerms().CheckNameExist(name); id > 0 {
		return errors.New(fmt.Sprintf("节点名称[%s]已经存在，请更换！", name))
	}
	return nil
}

func (s *Perms) Add(params dto.AddPerms) error {
	params.PagePerms = strings.ToUpper(params.PagePerms)
	permsModel := models.NewPerms()
	gconvert.StructCopy(params, permsModel)
	if err := NewPerms().checkPagePermsExist(permsModel.PagePerms); err != nil {
		return err
	}
	if err := NewPerms().checkNameExist(permsModel.Name); err != nil {
		return err
	}
	err := permsModel.Add()
	if err != nil {
		gsys.Logger.Error("添加节点失败—>", err.Error())
		return errors.New("添加节点失败！")
	}
	return nil
}

func (s *Perms) Update(params dto.UpdatePerms) error {
	params.PagePerms = strings.ToUpper(params.PagePerms)
	odlPermsModel := models.NewPerms()
	odlPermsModel.GetRow(params.Id)
	if odlPermsModel.Id == 0 {
		return errors.New("需要更新的节点不存在，请确认！")
	}
	//更新角色
	permsModel := models.NewPerms()
	gconvert.StructCopy(params, permsModel)
	if odlPermsModel.Name != permsModel.Name {
		if err := NewPerms().checkNameExist(permsModel.Name); err != nil {
			return err
		}
	}
	if odlPermsModel.PagePerms != permsModel.PagePerms {
		if err :=  NewPerms().checkPagePermsExist(permsModel.PagePerms); err != nil {
			return err
		}
	}
	err := permsModel.Update()
	if err != nil {
		gsys.Logger.Error("编辑节点失败—>", err.Error())
		return errors.New("更新节点失败！")
	}
	return nil
}