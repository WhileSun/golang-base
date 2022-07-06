package models

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/whilesun/go-admin/app/po"
	"github.com/whilesun/go-admin/app/vo"
	"github.com/whilesun/go-admin/pkg/utils/grequest"
)

type SPerms struct {
	po.SPerms
}

func NewPerms() *SPerms {
	return &SPerms{}
}

func (m *SPerms) GetList(req *gin.Context) ([]*SPerms, error) {
	bindParams := make(map[string]interface{}, 0)
	limit := grequest.PageLimit(req, bindParams)
	searchParams := map[string]string{"name": "like", "page_perms":"like"}
	where := grequest.ParamsWhere(req, searchParams, bindParams)
	perms := make([]*SPerms, 0)
	db.Raw(fmt.Sprintf("SELECT * FROM s_perms where 1=1 %s order by id desc %s",
		where, limit), bindParams).Scan(&perms)
	return perms, nil
}

func (m *SPerms) GetRow(id int) *SPerms {
	db.Model(&SPerms{}).Where("id = ?", id).First(m)
	return m
}

func (m *SPerms) Add() error {
	result := db.Create(m)
	return result.Error
}

func (m *SPerms) Update() error {
	return db.Model(m).Select("name", "page_perms", "data_perms").Updates(m).Error
}

func (m *SPerms) Delete() error{
	return db.Delete(m).Error
}

func (m *SPerms) CheckNameExist(name string) int {
	var id int
	db.Model(&SPerms{}).Select("id").Where("name = ?", name).Scan(&id)
	return id
}

func (m *SPerms) CheckPagePermsExist(pagePerms string) int{
	var id int
	db.Model(&SPerms{}).Select("id").Where("page_perms = ?", pagePerms).Scan(&id)
	return id
}

func (m *SPerms) GetFieldList() []*vo.PermsFieldList{
	perms := make([]*vo.PermsFieldList,0)
	db.Model(&SPerms{}).Select("id","name","page_perms").Order("id desc").Scan(&perms)
	return perms
}