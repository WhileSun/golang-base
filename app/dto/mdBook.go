package dto

type AddMdBook struct {
	BookName  string `form:"book_name"  binding:"required" label:"书籍名称"`
	BookIdent string `form:"book_ident"  binding:"" label:"书籍标识"`
}

type UpdateMdBook struct {
	Id       int    `form:"id"  binding:"required" label:"ID"`
	BookName string `form:"book_name"  binding:"required" label:"书籍名称"`
}
