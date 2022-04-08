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

type WorkTaskRecordController struct {
}

func (c *WorkTaskRecordController) GetList(req *gin.Context) {
	workTaskRecordModel := models.NewWorkTaskRecord()
	rows := workTaskRecordModel.GetList(req)
	e.New(req).Data(e.SUCCESS, rows)
}

func (c *WorkTaskRecordController) Add(req *gin.Context) {
	var params dto.AddWorkTaskRecord
	if err := gvalidator.ReqValidate(req, &params); err != nil {
		gsys.Logger.Info("添加工作项目任务参数有误->", err.Error())
		e.New(req).Msg(e.ERROR_API_PARAMS)
		return
	}
	workTaskRecordService := service.NewWorkTaskRecord()
	err := workTaskRecordService.Add(params)
	if err != nil {
		e.New(req).MsgDetail(e.FAILED, err.Error())
		return
	}
	e.New(req).Msg(e.SUCCESS)
}

func (c *WorkTaskRecordController) Update(req *gin.Context) {
	var params dto.UpdateWorkTaskRecord
	if err := gvalidator.ReqValidate(req, &params); err != nil {
		gsys.Logger.Info("修改工作项目任务参数有误->", err.Error())
		e.New(req).Msg(e.ERROR_API_PARAMS)
		return
	}
	workTaskRecordService := service.NewWorkTaskRecord()
	err := workTaskRecordService.Update(params)
	if err != nil {
		e.New(req).MsgDetail(e.FAILED, err.Error())
		return
	}
	e.New(req).Msg(e.SUCCESS)
}

func (c *WorkTaskRecordController) Delete(req *gin.Context){
	var params dto.DeleteWorkTaskRecord
	if err:= gvalidator.ReqValidate(req,&params);err!=nil{
		gsys.Logger.Info("删除工作项目任务参数有误->",err.Error())
		e.New(req).MsgDetail(e.ERROR_API_PARAMS,err.Error())
		return
	}
	workTaskRecordModel := models.NewWorkTaskRecord()
	workTaskRecordModel.Id = params.Id
	if err :=workTaskRecordModel.Delete();err !=nil{
		gsys.Logger.Error("删除工作项目任务失败—>", err.Error())
		e.New(req).MsgDetail(e.FAILED,"删除工作项目任务失败")
		return
	}
	e.New(req).Msg(e.SUCCESS)
}