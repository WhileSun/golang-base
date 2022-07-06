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

type Role struct {
	Base
}

func NewRole() *Role{
	return &Role{}
}

func (c *Role) Add(req *gin.Context) {
	var params dto.AddRole
	if err:= gvalidator.ReqValidate(req,&params);err!=nil{
		gsys.Logger.Info("添加角色参数有误->",err.Error())
		e.New(req).MsgDetail(e.ERROR_API_PARAMS,err.Error())
		return
	}
	roleService := service.NewRole()
	err := roleService.Add(params)
	if err != nil{
		e.New(req).MsgDetail(e.FAILED,err.Error())
		return
	}
	e.New(req).Msg(e.SUCCESS)
}

func (c *Role) Update(req *gin.Context){
	var params dto.UpdateRole
	if err:= gvalidator.ReqValidate(req,&params);err!=nil{
		gsys.Logger.Info("编辑角色参数有误->",err.Error())
		e.New(req).MsgDetail(e.ERROR_API_PARAMS,err.Error())
		return
	}
	roleService := service.NewRole()
	err := roleService.Update(params)
	if err != nil{
		e.New(req).MsgDetail(e.FAILED,err.Error())
		return
	}
	e.New(req).Msg(e.SUCCESS)
}

func (c *Role) GetList(req *gin.Context){
	roleModel := models.NewRole()
	rows, _ := roleModel.GetList(req)
	e.New(req).Data(e.SUCCESS, rows)
}


