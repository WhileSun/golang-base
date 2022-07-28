package models

import (
	"github.com/gin-gonic/gin"
	"github.com/whilesun/go-admin/app/po"
	"github.com/whilesun/go-admin/app/vo"
	"github.com/whilesun/go-admin/pkg/utils/gconvert"
	"gorm.io/gorm"
)

type MdDocument struct {
	po.MdDocument
}

func NewMdDocument() *MdDocument {
	return &MdDocument{}
}

func (m *MdDocument) CheckDocumentIdentExist(documentIdent string) int{
	var id int
	db.Model(&MdDocument{}).Select("id").Where("document_ident=?",documentIdent).Scan(&id)
	return id
}

func (m *MdDocument) GetNameList(req *gin.Context) []*vo.MdDocumentNameList{
	rows := make([]*vo.MdDocumentNameList,0)
	bookId := gconvert.StrToInt(req.PostForm("book_id"))
	db.Model(&MdDocument{}).Where("book_id=?",bookId).Order("order_sort asc,id asc").Find(&rows)
	return rows
}

func (m *MdDocument) AddName() error{
	return db.Create(m).Error
}

func (m *MdDocument) UpdateName() error{
	return db.Model(m).Select("document_name").Updates(m).Error
}

func (m *MdDocument) DeleteName(ids []int) error{
	return db.Where("id in ?",ids).Delete(m).Error
}

func (m *MdDocument) GetLastOrderSort(parentId int,bookId int) int{
	orderSort := 0
	db.Model(&MdDocument{}).Select("Max(order_sort) as order_sort").
		Where("parent_id = ? and book_id = ?",parentId,bookId).Scan(&orderSort)
	return orderSort
}

func (m *MdDocument) GetInfo(id int) *MdDocument{
	db.Model( &MdDocument{}).Where("id = ?",id).Find(m)
	return m
}

func (m *MdDocument) DragName(tx *gorm.DB,id int,parentId int,sort int) (err error){
	//移动后后面的栏目往后移一位
	err = tx.Model(&MdDocument{}).Where("parent_id = ? and order_sort >= ?",parentId,sort).
		Update("order_sort",gorm.Expr("order_sort+?",1)).Error
	if err != nil{
		return
	}
	//更新
	data := map[string]interface{}{"parent_id":parentId,"order_sort":sort}
	err = tx.Model(&MdDocument{}).Where("id = ?",id).Updates(data).Error
	return err
}

func (m *MdDocument) UpdateText(id int,mdText string,htmlText string) error{
	data:=map[string]interface{}{"md_text":mdText,"html_text":htmlText}
	return db.Model(&MdDocument{}).Where("id = ?",id).Updates(data).Error
}