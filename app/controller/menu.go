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

type MenuController struct {
	BaseController
}

func (c *MenuController) GetList(req *gin.Context){
	menuModel := models.NewMenu()
	rows, _ := menuModel.GetList()
	e.New(req).Data(e.SUCCESS, rows)
}

func (c *MenuController) Add(req *gin.Context) {
	var params dto.AddMenu
	if err:= gvalidator.ReqValidate(req,&params);err!=nil{
		gsys.Logger.Info("添加菜单参数有误->",err.Error())
		e.New(req).MsgDetail(e.ERROR_API_PARAMS,err.Error())
		return
	}
	menuService := service.NewMenu()
	err := menuService.Add(params)
	if err != nil{
		e.New(req).MsgDetail(e.FAILED,err.Error())
		return
	}
	e.New(req).Msg(e.SUCCESS)
}

func (c *MenuController) Update(req *gin.Context){
	var params dto.UpdateMenu
	if err:= gvalidator.ReqValidate(req,&params);err!=nil{
		gsys.Logger.Info("编辑菜单参数有误->",err.Error())
		e.New(req).MsgDetail(e.ERROR_API_PARAMS,err.Error())
		return
	}
	menuService := service.NewMenu()
	err := menuService.Update(params)
	if err != nil{
		e.New(req).MsgDetail(e.FAILED,err.Error())
		return
	}
	e.New(req).Msg(e.SUCCESS)
}

func (c *MenuController) Delete(req *gin.Context){
	var params dto.DeleteMenu
	if err:= gvalidator.ReqValidate(req,&params);err!=nil{
		gsys.Logger.Info("删除菜单参数有误->",err.Error())
		e.New(req).MsgDetail(e.ERROR_API_PARAMS,err.Error())
		return
	}
	menuService := service.NewMenu()
	err := menuService.Delete(params)
	if err != nil{
		e.New(req).MsgDetail(e.FAILED,err.Error())
		return
	}
	e.New(req).Msg(e.SUCCESS)
}


