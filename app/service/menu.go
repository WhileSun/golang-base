package service

import (
	"errors"
	"fmt"
	"github.com/whilesun/go-admin/app/dto"
	"github.com/whilesun/go-admin/app/models"
	"github.com/whilesun/go-admin/app/vo"
	"github.com/whilesun/go-admin/pkg/gsys"
	"github.com/whilesun/go-admin/pkg/utils/gconvert"
	"github.com/whilesun/go-admin/pkg/utils/gtools"
	"gorm.io/gorm"
	"strings"
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

func (s *MenuService) checkPagePermsExist(menuModel *models.SMenu) error {
	if id := menuModel.CheckPagePermsExist(); id > 0 {
		return errors.New(fmt.Sprintf("菜单操作权限标识[%s]已经存在，请更换！", menuModel.PagePerms))
	}
	return nil
}

func (s *MenuService) checkTypeParams(menu *models.SMenu) error {
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

func (s *MenuService) addPagePermsMenu(tx *gorm.DB, keys []string, parentId int, dataPermsHeader string) error {
	var perms []*vo.PermsMenuList
	tx.Model(models.NewPerms()).Select("name", "page_perms", "data_perms").Where("page_perms in ?", keys).Scan(&perms)
	existPerms := make([]string, 0)
	tx.Model(models.NewMenu()).Select("page_perms").Where("parent_id = ? and page_perms in ?", parentId, keys).Scan(&existPerms)
	if len(existPerms) > 0 {
		return errors.New(fmt.Sprintf("菜单操作权限标识[%s]已经存在，请更换！", strings.Join(existPerms, ",")))
	}
	menus := make([]*models.SMenu, 0)
	for key, perm := range perms {
		menu := &models.SMenu{}
		menu.MenuName = perm.Name
		menu.MenuType = 2
		menu.ParentId = parentId
		menu.Sort = key + 1
		menu.DataPerms = dataPermsHeader + perm.DataPerms
		menu.PagePerms = strings.ToUpper(perm.PagePerms)
		menus = append(menus, menu)
	}
	if err := tx.Create(menus).Error; err != nil {
		gsys.Logger.Error("添加菜单栏目子节点失败—>", err.Error())
		return errors.New("添加菜单栏目子节点失败！")
	}
	return nil
}

func (s *MenuService) Add(params dto.AddMenu) error {
	menuService := NewMenu()
	err := gsys.Db.Transaction(func(tx *gorm.DB) error {
		menuModel := models.NewMenu()
		//操作权限标识转大写
		params.PagePerms = strings.ToUpper(params.PagePerms)
		gconvert.StructCopy(params, menuModel)
		if err := menuService.checkTypeParams(menuModel); err != nil {
			return err
		}
		if menuModel.MenuType == 2 {
			if err := menuService.checkPagePermsExist(menuModel); err != nil {
				return err
			}
		}
		if err := tx.Create(menuModel).Error; err != nil {
			gsys.Logger.Error("添加菜单栏目失败—>", err.Error())
			return errors.New("添加菜单栏目失败！")
		}
		//只有菜单栏有节点
		if menuModel.MenuType == 1 && len(params.TargetKeys) > 0 {
			if err := menuService.addPagePermsMenu(tx, params.TargetKeys, menuModel.Id, params.DataPermsHeader); err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

func (s *MenuService) Update(params dto.UpdateMenu) error {
	menuService := NewMenu()
	err := gsys.Db.Transaction(func(tx *gorm.DB) error {
		//操作权限标识转大写
		params.PagePerms = strings.ToUpper(params.PagePerms)
		oldMenuModel := models.NewMenu()
		oldMenuModel.GetRow(params.Id)
		if oldMenuModel.Id == 0 {
			return errors.New("需要更新的菜单栏目不存在，请确认！")
		}
		if oldMenuModel.IsSys {
			return errors.New("系统默认栏目不支持修改！")
		}
		//生成参数model
		menuModel := models.NewMenu()
		gconvert.StructCopy(params, menuModel)
		if menuModel.MenuType == 1{
			addKeys := gtools.InArrayNotExist(params.TargetKeys, params.OldTargetKeys)
			if len(addKeys) > 0{
				if err := menuService.addPagePermsMenu(tx, addKeys, menuModel.Id, params.DataPermsHeader); err != nil {
					return err
				}
			}
			deleteKeys := gtools.InArrayNotExist(params.OldTargetKeys, params.TargetKeys)
			if len(deleteKeys)>0 {
				tx.Where("page_perms in ? and parent_id = ?", deleteKeys, menuModel.Id).Delete(models.NewMenu())
			}
		}else{
			if oldMenuModel.PagePerms != menuModel.PagePerms && oldMenuModel.MenuType==2{
				if err:=menuService.checkPagePermsExist(menuModel);err!=nil{
					return err
				}
			}
		}
		if err := menuModel.Update(tx); err != nil {
			gsys.Logger.Error("编辑菜单栏目失败—>", err.Error())
			return errors.New("编辑菜单栏目失败！")
		}
		return nil
	})
	return err
}

func (s *MenuService) Delete(params dto.DeleteMenu) error {
	oldMenuModel := models.NewMenu()
	oldMenuModel.GetRow(params.Id)
	if oldMenuModel.Id == 0 {
		return errors.New("需要删除的菜单栏目不存在，请确认！")
	}
	if oldMenuModel.IsSys {
		return errors.New("系统默认栏目不支持修改！")
	}
	total := oldMenuModel.CheckChildrenExists(params.Id)
	if total > 0 {
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
func (s *MenuService) GetUserRoutes(userId int) []*MenuRoutesResult {
	menuModel := models.NewMenu()
	isSuper, roles := NewUserAuth().GetRoles(userId)
	rows := menuModel.GetUserRoutesList(isSuper == 1, roles)
	menuRoutes := make([]*MenuRoutesResult, 0)
	menuRoutes = UserRoutesTrans(rows, 0, menuRoutes)
	return menuRoutes
}

func UserRoutesTrans(rows []*models.SMenu, parentId int, menuRoutes []*MenuRoutesResult) []*MenuRoutesResult {
	for _, row := range rows {
		if row.ParentId == parentId {
			if row.MenuType == 1 {
				tMenuRoutes := &MenuRoutesResult{}
				tMenuRoutes.Path = row.Url
				tMenuRoutes.Name = row.MenuName
				tMenuRoutes.Icon = row.Icon
				PagePerms := make([]string, 0)
				for _, subRow := range rows {
					if row.Id == subRow.ParentId && subRow.MenuType == 2 {
						PagePerms = append(PagePerms, subRow.PagePerms)
					}
				}
				tMenuRoutes.PagePerms = PagePerms
				subMenuRoutes := UserRoutesTrans(rows, row.Id, tMenuRoutes.Routes)
				//如果是目录，默认添加一个redirect
				if row.MenuType == 1 && len(subMenuRoutes) > 0 {
					tMenuRoutes.Routes = subMenuRoutes
					//tMenuRoutes.Routes = append(tMenuRoutes.Routes, &MenuRoutesResult{Path: row.Url, Redirect: subMenuRoutes[0].Path})
				}
				menuRoutes = append(menuRoutes, tMenuRoutes)
			}
		}
	}
	return menuRoutes
}
