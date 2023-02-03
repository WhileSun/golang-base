package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/whilesun/go-admin/app/api"
	"github.com/whilesun/go-admin/middleware"
)

func InitSysRouter(Router *gin.RouterGroup) {
	sysApi := api.ApiGroupApp.SysApiGroup
	{
		Router.GET("/sys/captcha/get", sysApi.SysGlobalApi.GetCaptcha)
		Router.POST("/user/login", sysApi.SysUserApi.Login)
		// 以上接口不需要登录
		Router.Use(middleware.LoginAuth())
		//sysRouter := Router.Group("/sys")
		//{
		//	sysRouter.
		//}
		userRouter := Router.Group("/user")
		{
			userRouter.POST("/out.login", sysApi.SysUserApi.OutLogin)
			userRouter.GET("/info/get", sysApi.SysUserApi.GetInfo)
			userRouter.POST("/list", sysApi.SysUserApi.GetList)
			userRouter.POST("/add", sysApi.SysUserApi.Add)
			userRouter.POST("/update", sysApi.SysUserApi.Update)
			userRouter.POST("/passwd/update", sysApi.SysUserApi.UpdatePasswd)
		}

		roleRoute := Router.Group("/role")
		{
			roleRoute.POST("/name/list", sysApi.SysRoleApi.GetNameList)
			roleRoute.POST("/add", sysApi.SysRoleApi.Add)
			roleRoute.POST("/update", sysApi.SysRoleApi.Update)
			roleRoute.POST("/list", sysApi.SysRoleApi.GetList)
		}

		menuRoute := Router.Group("/menu")
		{
			menuRoute.POST("/name/list", sysApi.SysMenuApi.GetNameList)
			menuRoute.POST("/add", sysApi.SysMenuApi.Add)
			menuRoute.POST("/update", sysApi.SysMenuApi.Update)
			menuRoute.POST("/list", sysApi.SysMenuApi.GetList)
			menuRoute.POST("/delete", sysApi.SysMenuApi.Delete)
		}

		permsRoute := Router.Group("/perms")
		{
			permsRoute.POST("/name/list", sysApi.SysPermsApi.GetNameList)
			permsRoute.POST("/add", sysApi.SysPermsApi.Add)
			permsRoute.POST("/update", sysApi.SysPermsApi.Update)
			permsRoute.POST("/list", sysApi.SysPermsApi.GetList)
			permsRoute.POST("/delete", sysApi.SysPermsApi.Delete)
		}
	}

}
