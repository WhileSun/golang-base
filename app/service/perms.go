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

type PermsService struct {
}

func NewPerms() *PermsService {
	return &PermsService{}
}

func checkPagePermsExist(permsModel *models.SPerms) error {
	if id := permsModel.CheckPagePermsExist(); id > 0 {
		return errors.New(fmt.Sprintf("节点操作权限标识[%s]已经存在，请更换！", permsModel.PagePerms))
	}
	return nil
}

func checkNameExist(permsModel *models.SPerms) error {
	if id := permsModel.CheckNameExist(); id > 0 {
		return errors.New(fmt.Sprintf("节点名称[%s]已经存在，请更换！", permsModel.Name))
	}
	return nil
}

func (s *PermsService) Add(params dto.AddPerms) error {
	params.PagePerms = strings.ToUpper(params.PagePerms)
	permsModel := models.NewPerms()
	gconvert.StructCopy(params, permsModel)
	if err := checkPagePermsExist(permsModel); err != nil {
		return err
	}
	if err := checkNameExist(permsModel); err != nil {
		return err
	}
	err := permsModel.Add()
	if err != nil {
		gsys.Logger.Error("添加节点失败—>", err.Error())
		return errors.New("添加节点失败！")
	}
	return nil
}

func (s *PermsService) Update(params dto.UpdatePerms) error {
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
		if err := checkNameExist(permsModel); err != nil {
			return err
		}
	}
	if odlPermsModel.PagePerms != permsModel.PagePerms {
		if err := checkPagePermsExist(permsModel); err != nil {
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