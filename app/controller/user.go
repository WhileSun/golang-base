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

type UserController struct {
	BaseController
}

//Login 用户登录接口
func (c *UserController) Login(req *gin.Context) {
	var params dto.LoginUser
	if err := gvalidator.ReqValidate(req, &params); err != nil {
		gsys.Logger.Info("用户登录参数有误->", err.Error())
		e.New(req).Msg(e.ERROR_API_PARAMS)
		return
	}
	userService := service.NewUser()
	token,code := userService.CheckLogin(&params)
	if code != e.SUCCESS {
		e.New(req).Msg(code)
		return
	}
	e.New(req).Data(e.SUCCESS, map[string]interface{}{"token": token})
}

//OutLogin 用户退出接口
func (c *UserController) OutLogin(req *gin.Context) {
	userId := req.GetInt("userId")
	userToken := req.GetString("userToken")
	userAuthService := service.NewUserAuth()
	userAuthService.DelSession(userId,userToken)
	e.New(req).Msg(e.SUCCESS)
}

func (c *UserController) GetList(req *gin.Context) {
	userModel := models.NewUser()
	rows := userModel.GetList(req)
	e.New(req).Data(e.SUCCESS, rows)
}

func (c *UserController) Add(req *gin.Context) {
	var params dto.AddUser
	if err := gvalidator.ReqValidate(req, &params); err != nil {
		gsys.Logger.Info("添加用户参数有误->", err.Error())
		e.New(req).Msg(e.ERROR_API_PARAMS)
		return
	}
	userService := service.NewUser()
	err := userService.Add(params)
	if err != nil {
		e.New(req).MsgDetail(e.FAILED, err.Error())
		return
	}
	e.New(req).Msg(e.SUCCESS)
}

func (c *UserController) Update(req *gin.Context) {
	var params dto.UpdateUser
	if err := gvalidator.ReqValidate(req, &params); err != nil {
		gsys.Logger.Info("修改用户参数有误->", err.Error())
		e.New(req).Msg(e.ERROR_API_PARAMS)
		return
	}
	userService := service.NewUser()
	err := userService.Update(params)
	if err != nil {
		e.New(req).MsgDetail(e.FAILED, err.Error())
		return
	}
	e.New(req).Msg(e.SUCCESS)
}

func (c *UserController) UpdatePasswd(req *gin.Context){
	var params dto.UpdateUserPasswd
	if err := gvalidator.ReqValidate(req, &params); err != nil {
		gsys.Logger.Info("修改用户密码参数有误->", err.Error())
		e.New(req).Msg(e.ERROR_API_PARAMS)
		return
	}
	userService := service.NewUser()
	err := userService.UpdatePasswd(params,req)
	if err != nil {
		e.New(req).MsgDetail(e.FAILED, err.Error())
		return
	}
	e.New(req).Msg(e.SUCCESS)
}

func (c *UserController) GetUserId(req *gin.Context) {
	userId := req.GetInt("userId")
	userModel := models.NewUser()
	userModel.GetInfo(userId)
	e.New(req).Data(e.SUCCESS, map[string]interface{}{"id": userModel.Id, "name": userModel.Realname, "avatar": "https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png"})
}

//GetUserRoutes 获取用户开放的路由权限
func (c *UserController) GetUserRoutes(req *gin.Context) {
	menuService := service.NewMenu()
	rows := menuService.GetUserRoutes(req.GetInt("userId"))
	e.New(req).Data(e.SUCCESS, rows)
}



