package models

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/whilesun/go-admin/app/po"
	"github.com/whilesun/go-admin/app/vo"
	"github.com/whilesun/go-admin/pkg/utils/grequest"
)

type WorkTaskProject struct {
	po.WorkTaskProject
}

func NewWorkTaskProject() *WorkTaskProject {
	return &WorkTaskProject{}
}

func (m *WorkTaskProject) CheckProjectNameExist() int{
	var id int
	db.Model(&WorkTaskProject{}).Select("id").Where("project_name=?",m.ProjectName).Scan(&id)
	return id
}

func (m *WorkTaskProject) GetRowById(id int){
	db.Where("id = ?",id).Find(m)
}

func (m *WorkTaskProject) GetList(req *gin.Context) []*WorkTaskProject{
	bindParams := make(map[string]interface{}, 0)
	limit := grequest.PageLimit(req, bindParams)
	searchParams := map[string]string{"project_name": "like"}
	where := grequest.ParamsWhere(req, searchParams, bindParams)
	projects := make([]*WorkTaskProject,0)
	db.Raw(fmt.Sprintf(`SELECT
	*
	FROM 
		work_task_project
		where 1=1 %s
		order by id desc %s`, where, limit), bindParams).Scan(&projects)
	return projects
}

func (m *WorkTaskProject) Add() error{
	return db.Create(m).Error
}

func (m *WorkTaskProject) Update() error{
	return db.Model(m).Select("project_name", "remark").Updates(m).Error
}

func (m *WorkTaskProject) GetFieldList() []*vo.WorkTaskProjectFieldList{
	projects := make([]*vo.WorkTaskProjectFieldList,0)
	db.Model(&WorkTaskProject{}).Select("id","project_name").Order("id desc").Scan(&projects)
	return projects
}