package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/whilesun/go-admin/app/models"
	"github.com/whilesun/go-admin/pkg/e"
)

type Load struct {

}

func NewLoad() *Load{
	return &Load{}
}

func (c *Load) GetRoleList(req *gin.Context){
	rows := models.NewRole().GetFieldList()
	e.New(req).Data(e.SUCCESS, rows)
}

func (c *Load) GetMenuList(req *gin.Context){
	rows := models.NewMenu().GetFieldList(req)
	e.New(req).Data(e.SUCCESS, rows)
}

func (c *Load) GetPermsList(req *gin.Context){
	rows := models.NewPerms().GetFieldList()
	e.New(req).Data(e.SUCCESS, rows)
}

func (c *Load) GetWorkProjectList(req *gin.Context){
	rows := models.NewWorkProject().GetFieldList()
	e.New(req).Data(e.SUCCESS, rows)
}
