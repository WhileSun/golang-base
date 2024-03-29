package service

import (
	"errors"
	"fmt"
	"github.com/whilesun/go-admin/app/dto"
	"github.com/whilesun/go-admin/app/models"
	"github.com/whilesun/go-admin/pkg/gsys"
	"github.com/whilesun/go-admin/pkg/utils/gconvert"
	"github.com/whilesun/go-admin/pkg/utils/gtools"
)

type MdBook struct {

}

func NewMdBook() *MdBook {
	return &MdBook{}
}

func (s *MdBook) checkBookIdentExist(bookIdent string) error {
	if id := models.NewMdBook().CheckBookIdentExist(bookIdent); id > 0 {
		return errors.New(fmt.Sprintf("书籍标识[%s]已经存在，请更换！", bookIdent))
	}
	return nil
}

func (s *MdBook) Add(params dto.AddMdBook) error{
	mdBookModel := models.NewMdBook()
	gconvert.StructCopy(params, mdBookModel)
	if mdBookModel.BookIdent == ""{
		mdBookModel.BookIdent = gtools.StrRand("book-")
	}else{
		if err := NewMdBook().checkBookIdentExist(mdBookModel.BookIdent); err != nil {
			return err
		}
	}
	if err := mdBookModel.Add();err !=nil{
		gsys.Logger.Error("添加书籍失败—>", err.Error())
		return errors.New("添加书籍失败！")
	}
	return nil
}

func (s *MdBook) Update(params dto.UpdateMdBook) error{
	mdBookModel := models.NewMdBook()
	gconvert.StructCopy(params, mdBookModel)
	if err :=mdBookModel.Update();err !=nil{
		gsys.Logger.Error("修改书籍失败—>", err.Error())
		return errors.New("修改书籍失败！")
	}
	return nil
}


