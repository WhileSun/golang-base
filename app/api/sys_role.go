package api

import (
	"github.com/gin-gonic/gin"
	"github.com/whilesun/go-admin/app/models"
	"github.com/whilesun/go-admin/app/types/dto"
	"github.com/whilesun/go-admin/gctx"
	"github.com/whilesun/go-admin/pkg/utils/e"
	"github.com/whilesun/go-admin/pkg/utils/gvalidator"
)

type SysRoleApi struct {
}

func (c *SysRoleApi) Add(req *gin.Context) {
	var params dto.AddRole
	if err := gvalidator.ReqValidate(req, &params); err != nil {
		gctx.Logger.Info("添加角色参数有误->", err.Error())
		e.New(req).MsgDetail(e.ERROR_API_PARAMS, err.Error())
		return
	}
	err := sysRoleService.Add(params)
	if err != nil {
		e.New(req).MsgDetail(e.FAILED, err.Error())
		return
	}
	e.New(req).Msg(e.SUCCESS)
}

func (c *SysRoleApi) Update(req *gin.Context) {
	var params dto.UpdateRole
	if err := gvalidator.ReqValidate(req, &params); err != nil {
		gctx.Logger.Info("编辑角色参数有误->", err.Error())
		e.New(req).MsgDetail(e.ERROR_API_PARAMS, err.Error())
		return
	}
	err := sysRoleService.Update(params)
	if err != nil {
		e.New(req).MsgDetail(e.FAILED, err.Error())
		return
	}
	e.New(req).Msg(e.SUCCESS)
}

func (c *SysRoleApi) GetList(req *gin.Context) {
	roleModel := models.NewRole()
	rows, _ := roleModel.GetList(req)
	e.New(req).Data(e.SUCCESS, rows)
}

// GetNameList 角色名称列表简单输出
func (c *SysRoleApi) GetNameList(req *gin.Context) {
	rows := models.NewRole().GetFieldList()
	e.New(req).Data(e.SUCCESS, rows)
}
