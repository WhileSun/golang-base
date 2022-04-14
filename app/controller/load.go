package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/whilesun/go-admin/app/models"
	"github.com/whilesun/go-admin/pkg/e"
)

type LoadController struct {

}

func NewLoad() *LoadController{
	return &LoadController{}
}

func (c *LoadController) GetRoleList(req *gin.Context){
	rows := models.NewRole().GetFieldList()
	e.New(req).Data(e.SUCCESS, rows)
}

func (c *LoadController) GetMenuList(req *gin.Context){
	rows := models.NewMenu().GetFieldList(req)
	e.New(req).Data(e.SUCCESS, rows)
}

func (c *LoadController) GetPermsList(req *gin.Context){
	rows := models.NewPerms().GetFieldList()
	e.New(req).Data(e.SUCCESS, rows)
}


func (c *LoadController) GetWorkTaskProjectList(req *gin.Context){
	rows := models.NewWorkTaskProject().GetFieldList()
	e.New(req).Data(e.SUCCESS, rows)
}
