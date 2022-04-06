package service

import (
	"errors"
	"fmt"
	"github.com/whilesun/go-admin/app/dto"
	"github.com/whilesun/go-admin/app/models"
	"github.com/whilesun/go-admin/pkg/gsys"
	"github.com/whilesun/go-admin/pkg/utils/gconvert"
)

type MenuService struct {
}

type MenuRoutesResult struct {
	Path       string              `json:"path,omitempty"`
	Icon       string              `json:"icon,omitempty"`
	Name       string              `json:"name,omitempty"`
	HideInMenu bool                `json:"hideInMenu,omitempty"`
	Redirect   string              `json:"redirect,omitempty"`
	PagePerms  []string            `json:"pagePerms,omitempty"`
	Routes     []*MenuRoutesResult `json:"routes,omitempty"`
}

func NewMenu() *MenuService {
	return &MenuService{}
}

func checkDataPermsExist(menuModel *models.SMenu) error {
	if id := menuModel.CheckDataPermsExist(); id > 0 {
		return errors.New(fmt.Sprintf("菜单数据权限标识[%s]已经存在，请更换！", menuModel.DataPerms))
	}
	return nil
}

func checkTypeParams(menu *models.SMenu) error {
	if menu.MenuType == 1 {
		if menu.Url == "" {
			return errors.New("菜单链接不能为空")
		}
	} else {
		if menu.DataPerms == "" {
			return errors.New("数据权限不能为空")
		}
	}
	return nil
}

func (s *MenuService) Add(params dto.AddMenu) error {
	menuModel := models.NewMenu()
	gconvert.StructCopy(params, menuModel)
	if err := checkTypeParams(menuModel); err != nil {
		return err
	}
	//if menuModel.MenuType==2{
	//	if err := checkDataPermsExist(menuModel); err != nil{
	//		return err
	//	}
	//}
	if err := menuModel.Add(); err != nil {
		gsys.Logger.Error("添加菜单栏目失败—>", err.Error())
		return errors.New("添加菜单栏目失败！")
	}
	return nil
}

func (s *MenuService) Update(params dto.UpdateMenu) error {
	oldMenuModel := models.NewMenu()
	oldMenuModel.GetRow(params.Id)
	if oldMenuModel.Id == 0 {
		return errors.New("需要更新的菜单栏目不存在，请确认！")
	}
	//生成参数model
	menuModel := models.NewMenu()
	gconvert.StructCopy(params, menuModel)
	//if oldMenuModel.DataPerms != menuModel.DataPerms && oldMenuModel.MenuType==4{
	//	if err:=checkDataPermsExist(menuModel);err!=nil{
	//		return err
	//	}
	//}
	err := menuModel.Update()
	if err != nil {
		gsys.Logger.Error("编辑菜单栏目失败—>", err.Error())
		return errors.New("编辑菜单栏目失败！")
	}
	return nil
}

func (s *MenuService) Delete(params dto.DeleteMenu) error{
	oldMenuModel := models.NewMenu()
	total := oldMenuModel.CheckChildrenExists(params.Id)
	if total>0{
		return errors.New("删除菜单目录失败，请先删除子目录！")
	}
	menuModel := models.NewMenu()
	menuModel.Id = params.Id
	err := menuModel.Delete()
	if err != nil {
		gsys.Logger.Error("删除菜单栏目失败—>", err.Error())
		return errors.New("删除菜单栏目失败！")
	}
	return nil
}


//GetUserRoutes 获取用户菜单栏权限
func (s *MenuService) GetUserRoutes(username string) []*MenuRoutesResult {
	menuModel := models.NewMenu()
	//super, permsIdArr := models.NewRolePolicy().GetUserRolesPerms(username)
	rows := menuModel.GetUserRoutesList(true, []string{})
	menuRoutes := make([]*MenuRoutesResult, 0)
	menuRoutes = UserRoutesTrans(rows, 0, menuRoutes)
	return menuRoutes
}

func UserRoutesTrans(rows []*models.SMenu, parentId int, menuRoutes []*MenuRoutesResult) []*MenuRoutesResult {
	for _, row := range rows {
		if row.ParentId == parentId {
			// MenuType 1,2为目录和菜单
			if row.MenuType < 3 {
				tMenuRoutes := &MenuRoutesResult{}
				tMenuRoutes.Path = row.Url
				tMenuRoutes.Name = row.MenuName
				tMenuRoutes.Icon = row.Icon
				if row.MenuType == 2 {
					PagePerms := make([]string, 0)
					for _, subRow := range rows {
						if row.Id == subRow.ParentId {
							PagePerms = append(PagePerms, subRow.PagePerms)
						}
					}
					tMenuRoutes.PagePerms = PagePerms
				}
				subMenuRoutes := UserRoutesTrans(rows, row.Id, tMenuRoutes.Routes)
				//如果是目录，默认添加一个redirect
				if row.MenuType == 1 && len(subMenuRoutes) > 0 {
					tMenuRoutes.Routes = subMenuRoutes
					tMenuRoutes.Routes = append(tMenuRoutes.Routes, &MenuRoutesResult{Path: row.Url, Redirect: subMenuRoutes[0].Path})
				}
				menuRoutes = append(menuRoutes, tMenuRoutes)
			}
		}
	}
	return menuRoutes
}
