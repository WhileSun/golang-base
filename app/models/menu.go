package models

import (
	"github.com/gin-gonic/gin"
	"github.com/whilesun/go-admin/app/po"
	"github.com/whilesun/go-admin/app/vo"
	"github.com/whilesun/go-admin/pkg/utils/gconvert"
)

type SMenu struct {
	po.SMenu
}

func NewMenu() *SMenu {
	return &SMenu{}
}

func (m *SMenu) GetList() ([]*SMenu, error) {
	var menus []*SMenu
	db.Order("sort asc").Find(&menus)
	return menus, nil
}

func (m *SMenu) Add() error {
	result := db.Create(m)
	return result.Error
}

func (m *SMenu) Update() error {
	result := db.Model(m).Select("menu_name","menu_type","url","icon","parent_id",
		"sort","data_perms","page_perms","status","show").Updates(m)
	return result.Error
}

func (m *SMenu) Delete() error {
	result := db.Delete(m)
	return result.Error
}

//CheckChildrenExists 判断下面是否有子集
func (m *SMenu) CheckChildrenExists(id int) int{
	var total int
	db.Model(&SMenu{}).Raw("SELECT count(id) as total FROM s_menu where parent_id = ?",id).Scan(&total)
	return total
}

//GetFieldList 获取菜单的字段列表
func (m *SMenu) GetFieldList(req *gin.Context) []*vo.MenuFieldList{
	menus := make([]*vo.MenuFieldList,0)
	menuType := gconvert.StrToInt(req.PostForm("menu_type"))
	con := db.Model(&SMenu{})
	if menuType>0 {
		con = con.Where("menu_type = ?",menuType)
	}
	con.Order("sort asc,id asc").Scan(&menus)
	return menus
}

func (m *SMenu) GetUserRoutesList(super bool, roles []string) []*SMenu{
	var menus []*SMenu
	if super{
		db.Model(&SMenu{}).Raw("SELECT * FROM s_menu where status=? order by sort asc",true).Scan(&menus)
	}else{
		menuIds := make([]string,0)
		db.Model(&SRole{}).Raw("select distinct regexp_split_to_table(perms_ids,',') as menu_ids from s_role where role_name in ? and status=?",roles,true).Scan(&menuIds)
		db.Model(&SMenu{}).Raw("SELECT * FROM s_menu where status=? and id in ? order by sort asc",true,menuIds).Scan(&menus)
	}
	return menus
}

func (m *SMenu) GeDataPerms(ids []string) []map[string]interface{}{
	resp := make([]map[string]interface{},0)
	db.Model(&SMenu{}).Select("data_perms").Where("id in ?",ids).
		Where("menu_type = ?",4).
		Where("data_perms <> ?","").Find(&resp)
	return resp
}


func (m *SMenu) GetRow(id int) *SMenu {
	db.Select("id,data_perms,status,menu_type").Where("id = ?", id).Find(m)
	return m
}

func (m *SMenu) CheckDataPermsExist() int{
	var id int
	db.Model(&SMenu{}).Select("id").Where("data_perms = ?", m.DataPerms).Scan(&id)
	return id
}