package models

import (
	"fmt"
	"github.com/gin-gonic/gin"
	po2 "github.com/whilesun/go-admin/app/types/po"
	vo2 "github.com/whilesun/go-admin/app/types/vo"
	"github.com/whilesun/go-admin/pkg/utils/grequest"
)

type WorkProject struct {
	po2.WorkProject
}

func NewWorkProject() *WorkProject {
	return &WorkProject{}
}

func (m *WorkProject) CheckProjectNameExist(projectName string) int {
	var id int
	db.Model(&WorkProject{}).Select("id").Where("project_name=?", projectName).Scan(&id)
	return id
}

func (m *WorkProject) GetInfo(id int) *WorkProject {
	db.Model(&WorkProject{}).Where("id = ?", id).First(m)
	return m
}

func (m *WorkProject) GetList(req *gin.Context) []*WorkProject {
	bindParams := make(map[string]interface{}, 0)
	limit := grequest.PageLimit(req, bindParams)
	searchParams := map[string]string{"project_name": "like"}
	where := grequest.ParamsWhere(req, searchParams, bindParams)
	projects := make([]*WorkProject, 0)
	db.Raw(fmt.Sprintf(`SELECT
	*
	FROM 
		work_project
		where 1=1 %s
		order by id desc %s`, where, limit), bindParams).Scan(&projects)
	return projects
}

func (m *WorkProject) Add() error {
	return db.Create(m).Error
}

func (m *WorkProject) Update() error {
	return db.Model(m).Select("project_name", "remark").Updates(m).Error
}

func (m *WorkProject) GetFieldList() []*vo2.WorkProjectFieldList {
	projects := make([]*vo2.WorkProjectFieldList, 0)
	db.Model(&WorkProject{}).Select("id", "project_name").Order("id desc").Scan(&projects)
	return projects
}
