package models

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/whilesun/go-admin/app/po"
	"github.com/whilesun/go-admin/pkg/utils/grequest"
)

type MdBook struct {
	po.MdBook
}

func NewMdBook() *MdBook {
	return &MdBook{}
}

func (m *MdBook) CheckBookIdentExist(bookIdent string) int{
	var id int
	db.Model(&MdBook{}).Select("id").Where("book_ident=?",bookIdent).Scan(&id)
	return id
}

func (m *MdBook) GetInfo(id int) *MdBook{
	db.Model( &MdBook{}).Where("id = ?",id).Find(m)
	return m
}

func (m *MdBook) GetList(req *gin.Context) []*MdBook{
	bindParams := make(map[string]interface{}, 0)
	limit := grequest.PageLimit(req, bindParams)
	searchParams := map[string]string{"book_name": "like"}
	where := grequest.ParamsWhere(req, searchParams, bindParams)
	rows := make([]*MdBook,0)
	db.Raw(fmt.Sprintf("SELECT * FROM md_book where 1=1 %s order by id desc %s",
		where, limit), bindParams).Scan(&rows)
	return rows
}

func (m *MdBook) Add() error{
	return db.Create(m).Error
}

func (m *MdBook) Update() error{
	return db.Model(m).Select("book_name").Updates(m).Error
}