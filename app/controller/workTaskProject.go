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

type WorkTaskProjectController struct {

}

func (c *WorkTaskProjectController) GetList(req *gin.Context){
	workTaskProjectModel := models.NewWorkTaskProject()
	rows := workTaskProjectModel.GetList(req)
	e.New(req).Data(e.SUCCESS, rows)

}

func (c *WorkTaskProjectController) Add(req *gin.Context){
	var params dto.AddWorkTaskProject
	if err := gvalidator.ReqValidate(req, &params); err != nil {
		gsys.Logger.Info("添加工作项目参数有误->", err.Error())
		e.New(req).Msg(e.ERROR_API_PARAMS)
		return
	}
	workTaskProjectService := service.NewWorkTaskProject()
	err := workTaskProjectService.Add(params,req)
	if err != nil {
		e.New(req).MsgDetail(e.FAILED, err.Error())
		return
	}
	e.New(req).Msg(e.SUCCESS)
}

func (c *WorkTaskProjectController) Update(req *gin.Context){
	var params dto.UpdateWorkTaskProject
	if err := gvalidator.ReqValidate(req, &params); err != nil {
		gsys.Logger.Info("修改工作项目参数有误->", err.Error())
		e.New(req).Msg(e.ERROR_API_PARAMS)
		return
	}
	workTaskProjectService := service.NewWorkTaskProject()
	err := workTaskProjectService.Update(params)
	if err != nil {
		e.New(req).MsgDetail(e.FAILED, err.Error())
		return
	}
	e.New(req).Msg(e.SUCCESS)
}