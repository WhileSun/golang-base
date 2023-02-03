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

type WorkProject struct {
}

func NewWorkProject() *WorkProject {
	return &WorkProject{}
}

func (c *WorkProject) GetList(req *gin.Context) {
	workProjectModel := models.NewWorkProject()
	rows := workProjectModel.GetList(req)
	e2.New(req).Data(e2.SUCCESS, rows)
}

func (c *WorkProject) Add(req *gin.Context) {
	var params dto2.AddWorkProject
	if err := gvalidator2.ReqValidate(req, &params); err != nil {
		gsys2.Logger.Info("添加工作项目参数有误->", err.Error())
		e2.New(req).Msg(e2.ERROR_API_PARAMS)
		return
	}
	workProjectService := service.NewWorkProject()
	err := workProjectService.Add(params, req)
	if err != nil {
		e2.New(req).MsgDetail(e2.FAILED, err.Error())
		return
	}
	e2.New(req).Msg(e2.SUCCESS)
}

func (c *WorkProject) Update(req *gin.Context) {
	var params dto2.UpdateWorkProject
	if err := gvalidator2.ReqValidate(req, &params); err != nil {
		gsys2.Logger.Info("修改工作项目参数有误->", err.Error())
		e2.New(req).Msg(e2.ERROR_API_PARAMS)
		return
	}
	workProjectService := service.NewWorkProject()
	err := workProjectService.Update(params)
	if err != nil {
		e2.New(req).MsgDetail(e2.FAILED, err.Error())
		return
	}
	e2.New(req).Msg(e2.SUCCESS)
}
