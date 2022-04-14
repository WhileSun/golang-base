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
	r.StaticFile("/", "dist/index.html") //前端接口

	r.POST("/test1", (&controller.IndexController{}).Test)

	//r.Use(middleware.Cors())
	r.Use(middleware.LoggerToFile())
	// 使用 Recovery 中间件
	r.Use(gin.Recovery())
	api := r.Group("/api")
	{
		api.GET("/sys/loginCaptcha/get", (&controller.SysController{}).GetLoginCaptcha)
		api.POST("/user/login", (&controller.UserController{}).Login)
		//以上接口不需要登录
		api.Use(middleware.LoginAuth())

		api.POST("/user/outLogin", (&controller.UserController{}).OutLogin)
		api.GET("/user/info/get", (&controller.UserController{}).GetUserId)
		api.POST("/user/routes/get", (&controller.UserController{}).GetUserRoutes)

		load := api.Group("/load")
		{
			load.POST("/role/list/get", (&controller.LoadController{}).GetRoleList)
			load.POST("/menu/list/get", (&controller.LoadController{}).GetMenuList)
			load.POST("/perms/list/get", (&controller.LoadController{}).GetPermsList)
			load.POST("/work/taskProject/list/get", (&controller.LoadController{}).GetWorkTaskProjectList)
		}
		//以上接口不需要接口权限配置
		api.Use(middleware.ReqAuth())

		user := api.Group("/user")
		{
			user.POST("/list/get", (&controller.UserController{}).GetList)
			user.POST("/add", (&controller.UserController{}).Add)
			user.POST("/update", (&controller.UserController{}).Update)
			user.POST("/passwd/update", (&controller.UserController{}).UpdatePasswd)
		}

		menu := api.Group("/menu")
		{
			menu.POST("/add", (&controller.MenuController{}).Add)
			menu.POST("/update", (&controller.MenuController{}).Update)
			menu.POST("/list/get", (&controller.MenuController{}).GetList)
			menu.POST("/delete", (&controller.MenuController{}).Delete)
		}

		perms := api.Group("/perms")
		{
			perms.POST("/add", (&controller.PermsController{}).Add)
			perms.POST("/update", (&controller.PermsController{}).Update)
			perms.POST("/list/get", (&controller.PermsController{}).GetList)
			perms.POST("/delete", (&controller.PermsController{}).Delete)
		}

		role := api.Group("/role")
		{
			role.POST("/add", (&controller.RoleController{}).Add)
			role.POST("/update", (&controller.RoleController{}).Update)
			role.POST("/list/get", (&controller.RoleController{}).GetList)
		}

		work := api.Group("/work")
		{
			work.POST("/taskProject/list/get", (&controller.WorkTaskProjectController{}).GetList)
			work.POST("/taskProject/add", (&controller.WorkTaskProjectController{}).Add)
			work.POST("/taskProject/update", (&controller.WorkTaskProjectController{}).Update)

			work.POST("/taskRecord/list/get", (&controller.WorkTaskRecordController{}).GetList)
			work.POST("/taskRecord/add", (&controller.WorkTaskRecordController{}).Add)
			work.POST("/taskRecord/update", (&controller.WorkTaskRecordController{}).Update)
			work.POST("/taskRecord/delete", (&controller.WorkTaskRecordController{}).Delete)
		}
	}
	return r
}
