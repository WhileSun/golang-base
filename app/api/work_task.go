package api

import (
	"github.com/gin-gonic/gin"
	"github.com/whilesun/go-admin/app/models"
	"github.com/whilesun/go-admin/app/service"
	dto2 "github.com/whilesun/go-admin/app/types/dto"
	gsys2 "github.com/whilesun/go-admin/gctx"
	e2 "github.com/whilesun/go-admin/pkg/utils/e"
	gvalidator2 "github.com/whilesun/go-admin/pkg/utils/gvalidator"
)

type WorkTask struct {
}

func NewWorkTask() *WorkTask {
	return &WorkTask{}
}

func (c *WorkTask) GetList(req *gin.Context) {
	workTaskModel := models.NewWorkTask()
	rows := workTaskModel.GetList(req)
	e2.New(req).Data(e2.SUCCESS, rows)
}

func (c *WorkTask) Add(req *gin.Context) {
	var params dto2.AddWorkTask
	if err := gvalidator2.ReqValidate(req, &params); err != nil {
		gsys2.Logger.Info("添加工作项目任务参数有误->", err.Error())
		e2.New(req).Msg(e2.ERROR_API_PARAMS)
		return
	}
	workTaskService := service.NewWorkTask()
	err := workTaskService.Add(params)
	if err != nil {
		e2.New(req).MsgDetail(e2.FAILED, err.Error())
		return
	}
	e2.New(req).Msg(e2.SUCCESS)
}

func (c *WorkTask) Update(req *gin.Context) {
	var params dto2.UpdateWorkTask
	if err := gvalidator2.ReqValidate(req, &params); err != nil {
		gsys2.Logger.Info("修改工作项目任务参数有误->", err.Error())
		e2.New(req).Msg(e2.ERROR_API_PARAMS)
		return
	}
	workTaskService := service.NewWorkTask()
	err := workTaskService.Update(params)
	if err != nil {
		e2.New(req).MsgDetail(e2.FAILED, err.Error())
		return
	}
	e2.New(req).Msg(e2.SUCCESS)
}

func (c *WorkTask) Delete(req *gin.Context) {
	var params dto2.DeleteWorkTask
	if err := gvalidator2.ReqValidate(req, &params); err != nil {
		gsys2.Logger.Info("删除工作项目任务参数有误->", err.Error())
		e2.New(req).MsgDetail(e2.ERROR_API_PARAMS, err.Error())
		return
	}
	workTaskModel := models.NewWorkTask()
	workTaskModel.Id = params.Id
	if err := workTaskModel.Delete(); err != nil {
		gsys2.Logger.Error("删除工作项目任务失败—>", err.Error())
		e2.New(req).MsgDetail(e2.FAILED, "删除工作项目任务失败")
		return
	}
	e2.New(req).Msg(e2.SUCCESS)
}

func (c *WorkTask) UploadPics(req *gin.Context) {
	common := NewCommon()
	urls, err := common.UploadPics(req, "work_task", "image")
	if err != nil {
		gsys2.Logger.Error("工作项目上传图片失败—>", err.Error())
		e2.New(req).MsgDetail(e2.FAILED, "上传图片失败")
		return
	}
	e2.New(req).Data(e2.SUCCESS, map[string]interface{}{"url": urls})
}
