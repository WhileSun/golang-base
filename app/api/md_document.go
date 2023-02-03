package api

import (
	"github.com/gin-gonic/gin"
	"github.com/whilesun/go-admin/app/models"
	"github.com/whilesun/go-admin/app/service"
	dto2 "github.com/whilesun/go-admin/app/types/dto"
	"github.com/whilesun/go-admin/gctx"
	e2 "github.com/whilesun/go-admin/pkg/utils/e"
	"github.com/whilesun/go-admin/pkg/utils/gconvert"
	gvalidator2 "github.com/whilesun/go-admin/pkg/utils/gvalidator"
)

type MdDocument struct {
}

func NewMdDocument() *MdDocument {
	return &MdDocument{}
}

func (c *MdDocument) GetNameList(req *gin.Context) {
	rows := models.NewMdDocument().GetNameList(req)
	info := models.NewMdBook().GetInfo(gconvert.StrToInt(req.PostForm("book_id")))
	e2.New(req).Data(e2.SUCCESS, map[string]interface{}{"book": info, "menuList": rows})
}

func (c *MdDocument) AddName(req *gin.Context) {
	var params dto2.AddMdDocumentName
	if err := gvalidator2.ReqValidate(req, &params); err != nil {
		gctx.Logger.Info("添加文档目录参数有误->", err.Error())
		e2.New(req).Msg(e2.ERROR_API_PARAMS)
		return
	}
	err := service.NewMdDocument().AddName(params)
	if err != nil {
		e2.New(req).MsgDetail(e2.FAILED, err.Error())
		return
	}
	e2.New(req).Msg(e2.SUCCESS)
}

func (c *MdDocument) UpdateName(req *gin.Context) {
	var params dto2.UpdateMdDocumentName
	if err := gvalidator2.ReqValidate(req, &params); err != nil {
		gctx.Logger.Info("修改文档目录参数有误->", err.Error())
		e2.New(req).Msg(e2.ERROR_API_PARAMS)
		return
	}
	err := service.NewMdDocument().UpdateName(params)
	if err != nil {
		e2.New(req).MsgDetail(e2.FAILED, err.Error())
		return
	}
	e2.New(req).Msg(e2.SUCCESS)
}

func (c *MdDocument) DeleteName(req *gin.Context) {
	var params dto2.DeleteMdDocumentName
	if err := gvalidator2.ReqValidate(req, &params); err != nil {
		gctx.Logger.Info("删除文档目录参数有误->", err.Error())
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

func (c *MdDocument) DragName(req *gin.Context) {
	var params dto2.DragMdDocumentName
	if err := gvalidator2.ReqValidate(req, &params); err != nil {
		gctx.Logger.Info("移动文档目录参数有误->", err.Error())
		e2.New(req).Msg(e2.ERROR_API_PARAMS)
		return
	}
	err := service.NewMdDocument().DragName(params)
	if err != nil {
		e2.New(req).MsgDetail(e2.FAILED, err.Error())
		return
	}
	e2.New(req).Msg(e2.SUCCESS)
}

func (c *MdDocument) GetText(req *gin.Context) {
	documentId := gconvert.StrToInt(req.PostForm("document_id"))
	if documentId == 0 {
		gctx.Logger.Info("获取文档内容参数有误")
		e2.New(req).Msg(e2.ERROR_API_PARAMS)
		return
	}
	row := models.NewMdDocument().GetInfo(documentId)
	e2.New(req).Data(e2.SUCCESS, row)
}

func (c *MdDocument) UpdateText(req *gin.Context) {
	var params dto2.UpdateMdDocumentText
	if err := gvalidator2.ReqValidate(req, &params); err != nil {
		gctx.Logger.Info("更新文档目录参数有误->", err.Error())
		e2.New(req).Msg(e2.ERROR_API_PARAMS)
		return
	}
	err := models.NewMdDocument().UpdateText(params.DocumentId, params.MdText, params.HtmlText)
	if err != nil {
		e2.New(req).MsgDetail(e2.FAILED, err.Error())
		return
	}
	e2.New(req).Msg(e2.SUCCESS)
}

func (c *MdDocument) UploadFile(req *gin.Context) {
	common := NewCommon()
	url, err := common.UploadFile(req, "md_document", "file")
	if err != nil {
		gctx.Logger.Error("上传失败—>", err.Error())
		e2.New(req).MsgDetail(e2.FAILED, "上传失败")
		return
	}
	e2.New(req).Data(e2.SUCCESS, map[string]interface{}{"url": url})
}
