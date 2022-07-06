package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/whilesun/go-admin/app/dto"
	"github.com/whilesun/go-admin/app/models"
	"github.com/whilesun/go-admin/app/service"
	"github.com/whilesun/go-admin/pkg/e"
	"github.com/whilesun/go-admin/pkg/gsys"
	"github.com/whilesun/go-admin/pkg/gvalidator"
	"github.com/whilesun/go-admin/pkg/utils/gtools"
)

type User struct {
	Base
}

func NewUser() *User{
	return &User{}
}

//Login 用户登录接口
func (c *User) Login(req *gin.Context) {
	var params dto.LoginUser
	if err := gvalidator.ReqValidate(req, &params); err != nil {
		gsys.Logger.Info("用户登录参数有误->", err.Error())
		e.New(req).Msg(e.ERROR_API_PARAMS)
		return
	}
	userService := service.NewUser()
	token,code := userService.CheckLogin(params)
	if code != e.SUCCESS {
		e.New(req).Msg(code)
		return
	}
	//生成cookie
	req.SetCookie("token", token, 0, "", "", false, true)
	e.New(req).Data(e.SUCCESS, map[string]interface{}{"token": token})
}

//OutLogin 用户退出接口
func (c *User) OutLogin(req *gin.Context) {
	userToken := req.GetString("userToken")
	service.NewUserAuth().DelSession(userToken)
	e.New(req).Msg(e.SUCCESS)
}

func (c *User) GetList(req *gin.Context) {
	userModel := models.NewUser()
	rows := userModel.GetList(req)
	e.New(req).Data(e.SUCCESS, rows)
}

func (c *User) Add(req *gin.Context) {
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

func (c *User) Update(req *gin.Context) {
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

func (c *User) UpdatePasswd(req *gin.Context){
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

func (c *User) GetInfo(req *gin.Context) {
	userId := req.GetInt("userId")
	userModel := models.NewUser().GetInfo(userId)
	e.New(req).Data(e.SUCCESS, map[string]interface{}{"id": userModel.Id, "name": userModel.Realname, "avatar": "https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png"})
}

//GetRouteList 获取用户开放的路由权限
func (c *User) GetRouteList(req *gin.Context) {
	menuService := service.NewMenu()
	rows := menuService.GetUserRouteList(gtools.GetUserRoleIdents(req))
	e.New(req).Data(e.SUCCESS, rows)
}



