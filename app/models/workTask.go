package models

import (
	"fmt"
	"github.com/gin-gonic/gin"
	po2 "github.com/whilesun/go-admin/app/types/po"
	vo2 "github.com/whilesun/go-admin/app/types/vo"
	"github.com/whilesun/go-admin/pkg/utils/grequest"
)

type WorkTask struct {
	po2.WorkTask
}

func NewWorkTask() *WorkTask {
	return &WorkTask{}
}

func (m *WorkTask) GetList(req *gin.Context) []*vo2.WorkTaskList {
	bindParams := make(map[string]interface{}, 0)
	limit := grequest.PageLimit(req, bindParams)
	searchParams := map[string]string{"xm*project_name": "like", "rw*title": "like"}
	where := grequest.ParamsWhere(req, searchParams, bindParams)
	taskRecords := make([]*vo2.WorkTaskList, 0)
	db.Raw(fmt.Sprintf(`SELECT
	rw.*,xm.project_name
	FROM 
		work_task rw
		left join work_project xm on rw.project_id = xm.id
		where 1=1 %s
		order by rw.launch_time desc %s`, where, limit), bindParams).Scan(&taskRecords)
	return taskRecords
}

func (m *WorkTask) Add() error {
	return db.Create(m).Error
}

func (m *WorkTask) Update() error {
	return db.Model(m).Select("*").Omit("created_at").Updates(m).Error
}

func (m *WorkTask) Delete() error {
	return db.Delete(m).Error
}
