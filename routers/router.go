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
		api.GET("/sys/loginCaptcha/get", controller.NewSys().GetLoginCaptcha)
		api.POST("/user/login", controller.NewUser().Login)
		//以上接口不需要登录
		api.Use(middleware.LoginAuth())

		api.POST("/user/out.login", controller.NewUser().OutLogin)
		api.GET("/user/info/get", controller.NewUser().GetInfo)
		api.GET("/user/route/list/get", controller.NewUser().GetRouteList)

		load := api.Group("/load")
		{
			load.POST("/role/list/get", controller.NewLoad().GetRoleList)
			load.POST("/menu/list/get", controller.NewLoad().GetMenuList)
			load.POST("/perms/list/get", controller.NewLoad().GetPermsList)
			load.POST("/work_project/list/get",controller.NewLoad().GetWorkProjectList)
		}

		//以上接口不需要接口权限配置
		api.Use(middleware.ReqAuth())
		user := api.Group("/user")
		{
			user.POST("/list/get", controller.NewUser().GetList)
			user.POST("/add", controller.NewUser().Add)
			user.POST("/update", controller.NewUser().Update)
			user.POST("/passwd/update", controller.NewUser().UpdatePasswd)
		}

		menu := api.Group("/menu")
		{
			menu.POST("/add", controller.NewMenu().Add)
			menu.POST("/update", controller.NewMenu().Update)
			menu.POST("/list/get", controller.NewMenu().GetList)
			menu.POST("/delete",controller.NewMenu().Delete)
		}

		perms := api.Group("/perms")
		{
			perms.POST("/add", controller.NewPerms().Add)
			perms.POST("/update", controller.NewPerms().Update)
			perms.POST("/list/get", controller.NewPerms().GetList)
			perms.POST("/delete", controller.NewPerms().Delete)
		}

		role := api.Group("/role")
		{
			role.POST("/add", controller.NewRole().Add)
			role.POST("/update", controller.NewRole().Update)
			role.POST("/list/get", controller.NewRole().GetList)
		}

		workProject := api.Group("/work_project")
		{
			workProject.POST("/list/get", controller.NewWorkProject().GetList)
			workProject.POST("/add", controller.NewWorkProject().Add)
			workProject.POST("/update", controller.NewWorkProject().Update)
		}

		workTask := api.Group("/work_task")
		{
			workTask.POST("/list/get", controller.NewWorkTask().GetList)
			workTask.POST("/add",  controller.NewWorkTask().Add)
			workTask.POST("/update",  controller.NewWorkTask().Update)
			workTask.POST("/delete",  controller.NewWorkTask().Delete)
			workTask.POST("/upload_pics", controller.NewWorkTask().UploadPics)
		}

		{
			api.POST("md_document_name/list/get",  controller.NewMdDocument().GetNameList)
			api.POST("md_document_name/add",  controller.NewMdDocument().AddName)
			api.POST("md_document_name/update",  controller.NewMdDocument().UpdateName)
			api.POST("md_document_name/delete",  controller.NewMdDocument().DeleteName)
			api.POST("md_document_name/drag",  controller.NewMdDocument().DragName)
			api.POST("md_document_text/get",  controller.NewMdDocument().GetText)
			api.POST("md_document_text/update",  controller.NewMdDocument().UpdateText)
		}

		{
			api.POST("md_book/list/get",  controller.NewMdBook().GetList)
			api.POST("md_book/add",  controller.NewMdBook().Add)
			api.POST("md_book/update",  controller.NewMdBook().Update)
		}
	}
	return r
}
