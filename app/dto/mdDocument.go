package dto

type AddMdDocumentName struct {
	DocumentName  string `form:"document_name"  binding:"required" label:"文档名称"`
	DocumentIdent string `form:"document_ident"  binding:"" label:"文档标识"`
	BookId        int    `form:"book_id" binding:"required,gt=0" label:"书籍ID"`
	ParentId      *int   `form:"parent_id" binding:"required,gte=0" label:"父级ID"`
}

type UpdateMdDocumentName struct {
	Id           int    `form:"id"  binding:"required" label:"ID"`
	DocumentName string `form:"document_name"  binding:"required" label:"文档名称"`
}

type DeleteMdDocumentName struct {
	Ids []int `form:"ids" binding:"required" label:"文档目录ID"`
}

type DragMdDocumentName struct {
	DragNodeId   int   `form:"drag_node_id"  binding:"required" label:"被drag的ID"`
	NodeId       int   `form:"node_id"  binding:"required" label:"drag的ID"`
	DragPosition *int  `form:"drag_position"  binding:"required" label:"移动的位置标志"`
	DragGap      *bool `form:"drag_gap"  binding:"required" label:"是否同级"`
}

type UpdateMdDocumentText struct {
	DocumentId int    `form:"document_id"  binding:"required" label:"ID"`
	MdText     string `form:"md_text"  binding:"required" label:"文档内容"`
	HtmlText   string `form:"html_text"  binding:"required" label:"文档内容"`
}
