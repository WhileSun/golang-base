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

type MdBook struct {

}

func NewMdBook() *MdBook{
	return &MdBook{}
}

func (c *MdBook) GetList(req *gin.Context){
	rows := models.NewMdBook().GetList(req)
	e.New(req).Data(e.SUCCESS, rows)
}

func (c *MdBook) Add(req *gin.Context){
	var params dto.AddMdBook
	if err := gvalidator.ReqValidate(req, &params); err != nil {
		gsys.Logger.Info("添加书籍参数有误->", err.Error())
		e.New(req).Msg(e.ERROR_API_PARAMS)
		return
	}
	err := service.NewMdBook().Add(params)
	if err != nil {
		e.New(req).MsgDetail(e.FAILED, err.Error())
		return
	}
	e.New(req).Msg(e.SUCCESS)
}

func (c *MdBook) Update(req *gin.Context){
	var params dto.UpdateMdBook
	if err := gvalidator.ReqValidate(req, &params); err != nil {
		gsys.Logger.Info("修改书籍参数有误->", err.Error())
		e.New(req).Msg(e.ERROR_API_PARAMS)
		return
	}
	err := service.NewMdBook().Update(params)
	if err != nil {
		e.New(req).MsgDetail(e.FAILED, err.Error())
		return
	}
	e.New(req).Msg(e.SUCCESS)
}

func (c *MdBook) Delete(req *gin.Context){
	var params dto.DeleteMdDocumentName
	if err := gvalidator.ReqValidate(req, &params); err != nil {
		gsys.Logger.Info("删除文档目录参数有误->", err.Error())
		e.New(req).Msg(e.ERROR_API_PARAMS)
		return
	}
	err := models.NewMdDocument().DeleteName(params.Ids)
	if err != nil {
		e.New(req).MsgDetail(e.FAILED, err.Error())
		return
	}
	e.New(req).Msg(e.SUCCESS)
}