package api

import (
	"github.com/gin-gonic/gin"
	"github.com/whilesun/go-admin/app/models"
	"github.com/whilesun/go-admin/app/types/dto"
	"github.com/whilesun/go-admin/gctx"
	"github.com/whilesun/go-admin/pkg/utils/e"
	"github.com/whilesun/go-admin/pkg/utils/gvalidator"
)

type SysMenuApi struct {
}

func (c *SysMenuApi) GetList(req *gin.Context) {
	menuModel := models.NewMenu()
	rows, _ := menuModel.GetList()
	e.New(req).Data(e.SUCCESS, rows)
}

// GetNameList 获取目录的一些基础字段
func (c *SysMenuApi) GetNameList(req *gin.Context) {
	rows := models.NewMenu().GetFieldList(req)
	e.New(req).Data(e.SUCCESS, rows)
}

func (c *SysMenuApi) Add(req *gin.Context) {
	var params dto.AddMenu
	if err := gvalidator.ReqValidate(req, &params); err != nil {
		gctx.Logger.Info("添加菜单参数有误->", err.Error())
		e.New(req).MsgDetail(e.ERROR_API_PARAMS, err.Error())
		return
	}
	err := sysMenuService.Add(params)
	if err != nil {
		e.New(req).MsgDetail(e.FAILED, err.Error())
		return
	}
	e.New(req).Msg(e.SUCCESS)
}

func (c *SysMenuApi) Update(req *gin.Context) {
	var params dto.UpdateMenu
	if err := gvalidator.ReqValidate(req, &params); err != nil {
		gctx.Logger.Info("编辑菜单参数有误->", err.Error())
		e.New(req).MsgDetail(e.ERROR_API_PARAMS, err.Error())
		return
	}
	err := sysMenuService.Update(params)
	if err != nil {
		e.New(req).MsgDetail(e.FAILED, err.Error())
		return
	}
	e.New(req).Msg(e.SUCCESS)
}

func (c *SysMenuApi) Delete(req *gin.Context) {
	var params dto.DeleteMenu
	if err := gvalidator.ReqValidate(req, &params); err != nil {
		gctx.Logger.Info("删除菜单参数有误->", err.Error())
		e.New(req).MsgDetail(e.ERROR_API_PARAMS, err.Error())
		return
	}
	err := sysMenuService.Delete(params)
	if err != nil {
		e.New(req).MsgDetail(e.FAILED, err.Error())
		return
	}
	e.New(req).Msg(e.SUCCESS)
}
