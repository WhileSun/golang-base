package service

import (
	"errors"
	"fmt"
	"github.com/whilesun/go-admin/app/models"
	dto2 "github.com/whilesun/go-admin/app/types/dto"
	gsys2 "github.com/whilesun/go-admin/gctx"
	"github.com/whilesun/go-admin/pkg/utils/gconvert"
	"github.com/whilesun/go-admin/pkg/utils/gtools"
	"gorm.io/gorm"
)

type mdDocument struct {
}

func NewMdDocument() *mdDocument {
	return &mdDocument{}
}

func (s *mdDocument) checkDocumentIdentExist(documentIdent string) error {
	if id := models.NewMdDocument().CheckDocumentIdentExist(documentIdent); id > 0 {
		return errors.New(fmt.Sprintf("文档标识[%s]已经存在，请更换！", documentIdent))
	}
	return nil
}

func (s *mdDocument) AddName(params dto2.AddMdDocumentName) error {
	bookModel := models.NewMdBook().GetInfo(params.BookId)
	if bookModel.Id == 0 {
		return errors.New("添加文档目录失败,书籍不存在！")
	}
	mdDocumentModel := models.NewMdDocument()
	gconvert.StructCopy(params, mdDocumentModel)
	if mdDocumentModel.DocumentIdent == "" {
		mdDocumentModel.DocumentIdent = gtools.StrRand("")
	} else {
		if err := NewMdDocument().checkDocumentIdentExist(mdDocumentModel.DocumentIdent); err != nil {
			return err
		}
	}
	mdDocumentModel.OrderSort = models.NewMdDocument().GetLastOrderSort(*params.ParentId, params.BookId) + 1
	if err := mdDocumentModel.AddName(); err != nil {
		gsys2.Logger.Error("添加文档目录失败—>", err.Error())
		return errors.New("添加文档目录失败！")
	}
	return nil
}

func (s *mdDocument) UpdateName(params dto2.UpdateMdDocumentName) error {
	mdDocumentModel := models.NewMdDocument()
	gconvert.StructCopy(params, mdDocumentModel)
	if err := mdDocumentModel.UpdateName(); err != nil {
		gsys2.Logger.Error("修改文档目录失败—>", err.Error())
		return errors.New("修改文档目录失败！")
	}
	return nil
}

func (s *mdDocument) DragName(params dto2.DragMdDocumentName) error {
	id := params.DragNodeId
	parentId := 0
	sort := 1
	if *params.DragGap {
		//移到顶级第一个
		if *params.DragPosition == -1 {
			parentId = 0
			sort = 1
		} else {
			//查询父级信息
			info := models.NewMdDocument().GetInfo(params.NodeId)
			parentId = info.ParentId
			sort = info.OrderSort + 1
		}
	} else {
		parentId = params.NodeId
		sort = 1
	}
	err := gsys2.Db.Transaction(func(tx *gorm.DB) error {
		err := models.NewMdDocument().DragName(tx, id, parentId, sort)
		return err
	})
	return err
}
