package api

import (
	"github.com/gin-gonic/gin"
	"github.com/whilesun/go-admin/app/models"
	"github.com/whilesun/go-admin/app/service"
	dto2 "github.com/whilesun/go-admin/app/types/dto"
	gsys2 "github.com/whilesun/go-admin/gctx"
	e2 "github.com/whilesun/go-admin/pkg/utils/e"
	gvalidator2 "github.com/whilesun/go-admin/pkg/utils/gvalidator"
)

type MdBook struct {
}

func NewMdBook() *MdBook {
	return &MdBook{}
}

func (c *MdBook) GetList(req *gin.Context) {
	rows := models.NewMdBook().GetList(req)
	e2.New(req).Data(e2.SUCCESS, rows)
}

func (c *MdBook) Add(req *gin.Context) {
	var params dto2.AddMdBook
	if err := gvalidator2.ReqValidate(req, &params); err != nil {
		gsys2.Logger.Info("添加书籍参数有误->", err.Error())
		e2.New(req).Msg(e2.ERROR_API_PARAMS)
		return
	}
	err := service.NewMdBook().Add(params)
	if err != nil {
		e2.New(req).MsgDetail(e2.FAILED, err.Error())
		return
	}
	e2.New(req).Msg(e2.SUCCESS)
}

func (c *MdBook) Update(req *gin.Context) {
	var params dto2.UpdateMdBook
	if err := gvalidator2.ReqValidate(req, &params); err != nil {
		gsys2.Logger.Info("修改书籍参数有误->", err.Error())
		e2.New(req).Msg(e2.ERROR_API_PARAMS)
		return
	}
	err := service.NewMdBook().Update(params)
	if err != nil {
		e2.New(req).MsgDetail(e2.FAILED, err.Error())
		return
	}
	e2.New(req).Msg(e2.SUCCESS)
}

func (c *MdBook) Delete(req *gin.Context) {
	var params dto2.DeleteMdDocumentName
	if err := gvalidator2.ReqValidate(req, &params); err != nil {
		gsys2.Logger.Info("删除文档目录参数有误->", err.Error())
		e2.New(req).Msg(e2.ERROR_API_PARAMS)
		return
	}
	err := models.NewMdDocument().DeleteName(params.Ids)
	if err != nil {
		e2.New(req).MsgDetail(e2.FAILED, err.Error())
		return
	}
	e2.New(req).Msg(e2.SUCCESS)
}
