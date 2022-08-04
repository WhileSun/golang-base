package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/whilesun/go-admin/app/dto"
	"github.com/whilesun/go-admin/app/models"
	"github.com/whilesun/go-admin/app/service"
	"github.com/whilesun/go-admin/pkg/e"
	"github.com/whilesun/go-admin/pkg/gsys"
	"github.com/whilesun/go-admin/pkg/gvalidator"
	"github.com/whilesun/go-admin/pkg/utils/gconvert"
)

type MdDocument struct {

}

func NewMdDocument() *MdDocument{
	return &MdDocument{}
}

func (c *MdDocument) GetNameList(req *gin.Context){
	rows := models.NewMdDocument().GetNameList(req)
	info := models.NewMdBook().GetInfo(gconvert.StrToInt(req.PostForm("book_id")))
	e.New(req).Data(e.SUCCESS,map[string]interface{}{"book":info,"menuList":rows})
}

func (c *MdDocument) AddName(req *gin.Context){
	var params dto.AddMdDocumentName
	if err := gvalidator.ReqValidate(req, &params); err != nil {
		gsys.Logger.Info("添加文档目录参数有误->", err.Error())
		e.New(req).Msg(e.ERROR_API_PARAMS)
		return
	}
	err := service.NewMdDocument().AddName(params)
	if err != nil {
		e.New(req).MsgDetail(e.FAILED, err.Error())
		return
	}
	e.New(req).Msg(e.SUCCESS)
}

func (c *MdDocument) UpdateName(req *gin.Context){
	var params dto.UpdateMdDocumentName
	if err := gvalidator.ReqValidate(req, &params); err != nil {
		gsys.Logger.Info("修改文档目录参数有误->", err.Error())
		e.New(req).Msg(e.ERROR_API_PARAMS)
		return
	}
	err := service.NewMdDocument().UpdateName(params)
	if err != nil {
		e.New(req).MsgDetail(e.FAILED, err.Error())
		return
	}
	e.New(req).Msg(e.SUCCESS)
}

func (c *MdDocument) DeleteName(req *gin.Context){
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

func (c *MdDocument) DragName(req *gin.Context){
	var params dto.DragMdDocumentName
	if err := gvalidator.ReqValidate(req, &params); err != nil {
		gsys.Logger.Info("移动文档目录参数有误->", err.Error())
		e.New(req).Msg(e.ERROR_API_PARAMS)
		return
	}
	err := service.NewMdDocument().DragName(params)
	if err != nil {
		e.New(req).MsgDetail(e.FAILED, err.Error())
		return
	}
	e.New(req).Msg(e.SUCCESS)
}

func (c *MdDocument) GetText(req *gin.Context){
	documentId := gconvert.StrToInt( req.PostForm("document_id"))
	if documentId==0{
		gsys.Logger.Info("获取文档内容参数有误")
		e.New(req).Msg(e.ERROR_API_PARAMS)
		return
	}
	row := models.NewMdDocument().GetInfo(documentId)
	e.New(req).Data(e.SUCCESS, row)
}

func (c *MdDocument) UpdateText(req *gin.Context){
	var params dto.UpdateMdDocumentText
	if err := gvalidator.ReqValidate(req, &params); err != nil {
		gsys.Logger.Info("更新文档目录参数有误->", err.Error())
		e.New(req).Msg(e.ERROR_API_PARAMS)
		return
	}
	err := models.NewMdDocument().UpdateText(params.DocumentId,params.MdText,params.HtmlText)
	if err != nil {
		e.New(req).MsgDetail(e.FAILED, err.Error())
		return
	}
	e.New(req).Msg(e.SUCCESS)
}

func (c *MdDocument) UploadFile(req *gin.Context){
	common := NewCommon()
	url,err := common.UploadFile(req,"md_document","file")
	if err !=nil{
		gsys.Logger.Error("上传失败—>", err.Error())
		e.New(req).MsgDetail(e.FAILED,"上传失败")
		return
	}
	e.New(req).Data(e.SUCCESS,map[string]interface{}{"url":url})
}