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

type PermsController struct {
}


func (c *PermsController) GetList(req *gin.Context){
	permsModel := models.NewPerms()
	rows, _ := permsModel.GetList(req)
	e.New(req).Data(e.SUCCESS, rows)
}

func (c *PermsController) Add(req *gin.Context) {
	var params dto.AddPerms
	if err:= gvalidator.ReqValidate(req,&params);err!=nil{
		gsys.Logger.Info("添加节点参数有误->",err.Error())
		e.New(req).MsgDetail(e.ERROR_API_PARAMS,err.Error())
		return
	}
	permsService := service.NewPerms()
	err := permsService.Add(params)
	if err != nil{
		e.New(req).MsgDetail(e.FAILED,err.Error())
		return
	}
	e.New(req).Msg(e.SUCCESS)
}

func (c *PermsController) Update(req *gin.Context){
	var params dto.UpdatePerms
	if err:= gvalidator.ReqValidate(req,&params);err!=nil{
		gsys.Logger.Info("编辑节点参数有误->",err.Error())
		e.New(req).MsgDetail(e.ERROR_API_PARAMS,err.Error())
		return
	}
	permsService := service.NewPerms()
	err := permsService.Update(params)
	if err != nil{
		e.New(req).MsgDetail(e.FAILED,err.Error())
		return
	}
	e.New(req).Msg(e.SUCCESS)
}

func (c *PermsController) Delete(req *gin.Context){
	var params dto.DeletePerms
	if err:= gvalidator.ReqValidate(req,&params);err!=nil{
		gsys.Logger.Info("删除节点参数有误->",err.Error())
		e.New(req).MsgDetail(e.ERROR_API_PARAMS,err.Error())
		return
	}
	permsModel := models.NewPerms()
	permsModel.Id = params.Id
	if err :=permsModel.Delete();err !=nil{
		gsys.Logger.Error("删除节点失败—>", err.Error())
		e.New(req).MsgDetail(e.FAILED,"删除节点失败")
		return
	}
	e.New(req).Msg(e.SUCCESS)
}