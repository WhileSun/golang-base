package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/whilesun/go-admin/app/controller"
	"github.com/whilesun/go-admin/middleware"
)

func InitRouter() *gin.Engine {

	r := gin.New()
	r.Static("/static", "./dist/static")
	r.Static("/uploads", "./uploads")
	r.StaticFile("/", "dist/index.html")  //前端接口

	r.POST("/test1",(&controller.IndexController{}).Test)

	//r.Use(middleware.Cors())
	r.Use(middleware.LoggerToFile())
	// 使用 Recovery 中间件
	r.Use(gin.Recovery())
	api:= r.Group("/api")
	{
		api.GET("/sys/loginCaptcha/get", (&controller.SysController{}).GetLoginCaptcha)
		api.POST("/user/login", (&controller.UserController{}).Login)
		//以上接口不需要登录
		api.Use(middleware.LoginAuth())

		api.POST("/user/outLogin", (&controller.UserController{}).OutLogin)
		api.POST("/user/passwd/update", (&controller.UserController{}).UpdatePasswd)
		api.GET("/user/info/get", (&controller.UserController{}).GetUserId)
		api.POST("/user/routes/get", (&controller.UserController{}).GetUserRoutes)

		api.POST("load/role/list/get",(&controller.LoadController{}).GetRoleList)
		api.POST("load/menu/list/get",(&controller.LoadController{}).GetMenuList)

		//以上接口不需要接口权限配置
		api.Use(middleware.ReqAuth())

		api.POST("/user/list/get",(&controller.UserController{}).GetList)
		api.POST("/user/add",(&controller.UserController{}).Add)
		api.POST("/user/update",(&controller.UserController{}).Update)

		menu := api.Group("/menu")
		{
			menu.POST("/add",(&controller.MenuController{}).Add)
			menu.POST("/update",(&controller.MenuController{}).Update)
			menu.POST("/list/get",(&controller.MenuController{}).GetList)
			menu.POST("/delete",(&controller.MenuController{}).Delete)
		}

		role := api.Group("/role")
		{
			role.POST("/add",(&controller.RoleController{}).Add)
			role.POST("/update",(&controller.RoleController{}).Update)
			role.POST("/list/get",(&controller.RoleController{}).GetList)
		}
	}
	return r
}
