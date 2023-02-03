package api

import (
	"github.com/gin-gonic/gin"
	"github.com/whilesun/go-admin/app/models"
	"github.com/whilesun/go-admin/app/types/dto"
	"github.com/whilesun/go-admin/gctx"
	"github.com/whilesun/go-admin/pkg/utils/e"
	"github.com/whilesun/go-admin/pkg/utils/gtools"
	"github.com/whilesun/go-admin/pkg/utils/gvalidator"
)

type SysUserApi struct {
}

func (c *SysUserApi) GetList(req *gin.Context) {
	userModel := models.NewUser()
	rows := userModel.GetList(req)
	e.New(req).Data(e.SUCCESS, rows)
}

// Login 用户登录接口
func (c *SysUserApi) Login(req *gin.Context) {
	var params dto.LoginUser
	if err := gvalidator.ReqValidate(req, &params); err != nil {
		gctx.Logger.Info("用户登录参数有误->", err.Error())
		e.New(req).Msg(e.ERROR_API_PARAMS)
		return
	}
	token, code := sysUserService.CheckLogin(params)
	if code != e.SUCCESS {
		e.New(req).Msg(code)
		return
	}
	e.New(req).Data(e.SUCCESS, map[string]interface{}{"token": token})
}

// OutLogin 用户退出接口
func (c *SysUserApi) OutLogin(req *gin.Context) {
	e.New(req).Msg(e.SUCCESS)
}

func (c *SysUserApi) Add(req *gin.Context) {
	var params dto.AddUser
	if err := gvalidator.ReqValidate(req, &params); err != nil {
		gctx.Logger.Info("添加用户参数有误->", err.Error())
		e.New(req).Msg(e.ERROR_API_PARAMS)
		return
	}
	err := sysUserService.Add(params)
	if err != nil {
		e.New(req).MsgDetail(e.FAILED, err.Error())
		return
	}
	e.New(req).Msg(e.SUCCESS)
}

func (c *SysUserApi) Update(req *gin.Context) {
	var params dto.UpdateUser
	if err := gvalidator.ReqValidate(req, &params); err != nil {
		gctx.Logger.Info("修改用户参数有误->", err.Error())
		e.New(req).Msg(e.ERROR_API_PARAMS)
		return
	}
	err := sysUserService.Update(params)
	if err != nil {
		e.New(req).MsgDetail(e.FAILED, err.Error())
		return
	}
	e.New(req).Msg(e.SUCCESS)
}

// UpdatePasswd 修改密码
func (c *SysUserApi) UpdatePasswd(req *gin.Context) {
	var params dto.UpdateUserPasswd
	if err := gvalidator.ReqValidate(req, &params); err != nil {
		gctx.Logger.Info("修改用户密码参数有误->", err.Error())
		e.New(req).Msg(e.ERROR_API_PARAMS)
		return
	}
	err := sysUserService.UpdatePasswd(params, req)
	if err != nil {
		e.New(req).MsgDetail(e.FAILED, err.Error())
		return
	}
	e.New(req).Msg(e.SUCCESS)
}

func (c *SysUserApi) GetInfo(req *gin.Context) {
	userId := req.GetInt("userId")
	userModel := models.NewUser().GetInfo(userId)
	e.New(req).Data(e.SUCCESS, map[string]interface{}{"id": userModel.Id, "name": userModel.Realname, "avatar": "https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png"})
}

// GetRouteList 获取用户开放的路由权限
func (c *SysUserApi) GetRouteList(req *gin.Context) {
	rows := sysMenuService.GetUserRouteList(gtools.GetUserRoleIdents(req))
	e.New(req).Data(e.SUCCESS, rows)
}
