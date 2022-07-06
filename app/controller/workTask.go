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

type WorkTask struct {
}

func  NewWorkTask() *WorkTask{
	return &WorkTask{}
}

func (c *WorkTask) GetList(req *gin.Context) {
	workTaskModel := models.NewWorkTask()
	rows := workTaskModel.GetList(req)
	e.New(req).Data(e.SUCCESS, rows)
}

func (c *WorkTask) Add(req *gin.Context) {
	var params dto.AddWorkTask
	if err := gvalidator.ReqValidate(req, &params); err != nil {
		gsys.Logger.Info("添加工作项目任务参数有误->", err.Error())
		e.New(req).Msg(e.ERROR_API_PARAMS)
		return
	}
	workTaskService := service.NewWorkTask()
	err := workTaskService.Add(params)
	if err != nil {
		e.New(req).MsgDetail(e.FAILED, err.Error())
		return
	}
	e.New(req).Msg(e.SUCCESS)
}

func (c *WorkTask) Update(req *gin.Context) {
	var params dto.UpdateWorkTask
	if err := gvalidator.ReqValidate(req, &params); err != nil {
		gsys.Logger.Info("修改工作项目任务参数有误->", err.Error())
		e.New(req).Msg(e.ERROR_API_PARAMS)
		return
	}
	workTaskService := service.NewWorkTask()
	err := workTaskService.Update(params)
	if err != nil {
		e.New(req).MsgDetail(e.FAILED, err.Error())
		return
	}
	e.New(req).Msg(e.SUCCESS)
}

func (c *WorkTask) Delete(req *gin.Context){
	var params dto.DeleteWorkTask
	if err:= gvalidator.ReqValidate(req,&params);err!=nil{
		gsys.Logger.Info("删除工作项目任务参数有误->",err.Error())
		e.New(req).MsgDetail(e.ERROR_API_PARAMS,err.Error())
		return
	}
	workTaskModel := models.NewWorkTask()
	workTaskModel.Id = params.Id
	if err :=workTaskModel.Delete();err !=nil{
		gsys.Logger.Error("删除工作项目任务失败—>", err.Error())
		e.New(req).MsgDetail(e.FAILED,"删除工作项目任务失败")
		return
	}
	e.New(req).Msg(e.SUCCESS)
}

func (c *WorkTask) UploadPics(req *gin.Context){
	common := NewCommon()
	urls,err := common.UploadPics(req,"work_task","image")
	if err !=nil{
		gsys.Logger.Error("工作项目上传图片失败—>", err.Error())
		e.New(req).MsgDetail(e.FAILED,"上传图片失败")
		return
	}
	e.New(req).Data(e.SUCCESS,map[string]interface{}{"url":urls})
}