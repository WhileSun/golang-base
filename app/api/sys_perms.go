package api

import (
	"github.com/gin-gonic/gin"
	"github.com/whilesun/go-admin/app/models"
	dto2 "github.com/whilesun/go-admin/app/types/dto"
	gsys2 "github.com/whilesun/go-admin/gctx"
	e2 "github.com/whilesun/go-admin/pkg/utils/e"
	gvalidator2 "github.com/whilesun/go-admin/pkg/utils/gvalidator"
)

type SysPermsApi struct {
}

func (c *SysPermsApi) GetList(req *gin.Context) {
	permsModel := models.NewPerms()
	rows, _ := permsModel.GetList(req)
	e2.New(req).Data(e2.SUCCESS, rows)
}

func (c *SysPermsApi) GetNameList(req *gin.Context) {
	rows := models.NewPerms().GetFieldList()
	e2.New(req).Data(e2.SUCCESS, rows)
}

func (c *SysPermsApi) Add(req *gin.Context) {
	var params dto2.AddPerms
	if err := gvalidator2.ReqValidate(req, &params); err != nil {
		gsys2.Logger.Info("添加节点参数有误->", err.Error())
		e2.New(req).MsgDetail(e2.ERROR_API_PARAMS, err.Error())
		return
	}
	err := sysPermsService.Add(params)
	if err != nil {
		e2.New(req).MsgDetail(e2.FAILED, err.Error())
		return
	}
	e2.New(req).Msg(e2.SUCCESS)
}

func (c *SysPermsApi) Update(req *gin.Context) {
	var params dto2.UpdatePerms
	if err := gvalidator2.ReqValidate(req, &params); err != nil {
		gsys2.Logger.Info("编辑节点参数有误->", err.Error())
		e2.New(req).MsgDetail(e2.ERROR_API_PARAMS, err.Error())
		return
	}
	err := sysPermsService.Update(params)
	if err != nil {
		e2.New(req).MsgDetail(e2.FAILED, err.Error())
		return
	}
	e2.New(req).Msg(e2.SUCCESS)
}

func (c *SysPermsApi) Delete(req *gin.Context) {
	var params dto2.DeletePerms
	if err := gvalidator2.ReqValidate(req, &params); err != nil {
		gsys2.Logger.Info("删除节点参数有误->", err.Error())
		e2.New(req).MsgDetail(e2.ERROR_API_PARAMS, err.Error())
		return
	}
	permsModel := models.NewPerms()
	permsModel.Id = params.Id
	if err := permsModel.Delete(); err != nil {
		gsys2.Logger.Error("删除节点失败—>", err.Error())
		e2.New(req).MsgDetail(e2.FAILED, "删除节点失败")
		return
	}
	e2.New(req).Msg(e2.SUCCESS)
}
