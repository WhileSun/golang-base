package api

import (
	"github.com/gin-gonic/gin"
	"github.com/whilesun/go-admin/app/models"
	e2 "github.com/whilesun/go-admin/pkg/utils/e"
)

type Load struct {
}

func NewLoad() *Load {
	return &Load{}
}

func (c *Load) GetWorkProjectList(req *gin.Context) {
	rows := models.NewWorkProject().GetFieldList()
	e2.New(req).Data(e2.SUCCESS, rows)
}
