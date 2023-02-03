package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/whilesun/go-admin/middleware"
)

func InitRouter() *gin.Engine {

	r := gin.New()
	r.Static("/static", "./dist/static")
	r.Static("/uploads", "./uploads")
	r.StaticFile("/", "dist/index.html") //前端接口

	//r.Use(middleware.Cors())
	r.Use(middleware.LoggerToFile())
	// 使用 Recovery 中间件
	r.Use(gin.Recovery())
	api := r.Group("/api")
	{
		InitSysRouter(api)
		//api.GET("/user/route/list/get", api.NewUser().GetRouteList)
		//
		//load := api.Group("/load")
		//{
		//	load.POST("/role/list/get", cront)
		//	load.POST("/menu/list/get", api.NewLoad().GetMenuList)
		//	load.POST("/perms/list/get", api.NewLoad().GetPermsList)
		//	load.POST("/work_project/list/get", api.NewLoad().GetWorkProjectList)
		//}
		//
		////以上接口不需要接口权限配置
		//api.Use(middleware.ReqAuth())
		//
		//
		//workProject := api.Group("/work_project")
		//{
		//	workProject.POST("/list/get", api.NewWorkProject().GetList)
		//	workProject.POST("/add", api.NewWorkProject().Add)
		//	workProject.POST("/update", api.NewWorkProject().Update)
		//}
		//
		//workTask := api.Group("/work_task")
		//{
		//	workTask.POST("/list/get", api.NewWorkTask().GetList)
		//	workTask.POST("/add", api.NewWorkTask().Add)
		//	workTask.POST("/update", api.NewWorkTask().Update)
		//	workTask.POST("/delete", api.NewWorkTask().Delete)
		//	workTask.POST("/upload_pics", api.NewWorkTask().UploadPics)
		//}
		//
		//{
		//	api.POST("md_document_name/list/get", api.NewMdDocument().GetNameList)
		//	api.POST("md_document_name/add", api.NewMdDocument().AddName)
		//	api.POST("md_document_name/update", api.NewMdDocument().UpdateName)
		//	api.POST("md_document_name/delete", api.NewMdDocument().DeleteName)
		//	api.POST("md_document_name/drag", api.NewMdDocument().DragName)
		//	api.POST("md_document_name/upload_file", api.NewMdDocument().UploadFile)
		//	api.POST("md_document_text/get", api.NewMdDocument().GetText)
		//	api.POST("md_document_text/update", api.NewMdDocument().UpdateText)
		//}
		//
		//{
		//	api.POST("md_book/list/get", api.NewMdBook().GetList)
		//	api.POST("md_book/add", api.NewMdBook().Add)
		//	api.POST("md_book/update", api.NewMdBook().Update)
		//}
	}
	return r
}
