package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/whilesun/go-admin/app/dto"
	"github.com/whilesun/go-admin/app/models"
	"github.com/whilesun/go-admin/app/service"
	"github.com/whilesun/go-admin/pkg/e"
	"github.com/whilesun/go-admin/pkg/gsys"
	"github.com/whilesun/go-admin/pkg/gvalidator"
)

type WorkProject struct {

}

func NewWorkProject() *WorkProject{
	return &WorkProject{}
}

func (c *WorkProject) GetList(req *gin.Context){
	workProjectModel := models.NewWorkProject()
	rows := workProjectModel.GetList(req)
	e.New(req).Data(e.SUCCESS, rows)
}

func (c *WorkProject) Add(req *gin.Context){
	var params dto.AddWorkProject
	if err := gvalidator.ReqValidate(req, &params); err != nil {
		gsys.Logger.Info("添加工作项目参数有误->", err.Error())
		e.New(req).Msg(e.ERROR_API_PARAMS)
		return
	}
	workProjectService := service.NewWorkProject()
	err := workProjectService.Add(params,req)
	if err != nil {
		e.New(req).MsgDetail(e.FAILED, err.Error())
		return
	}
	e.New(req).Msg(e.SUCCESS)
}

func (c *WorkProject) Update(req *gin.Context){
	var params dto.UpdateWorkProject
	if err := gvalidator.ReqValidate(req, &params); err != nil {
		gsys.Logger.Info("修改工作项目参数有误->", err.Error())
		e.New(req).Msg(e.ERROR_API_PARAMS)
		return
	}
	workProjectService := service.NewWorkProject()
	err := workProjectService.Update(params)
	if err != nil {
		e.New(req).MsgDetail(e.FAILED, err.Error())
		return
	}
	e.New(req).Msg(e.SUCCESS)
}