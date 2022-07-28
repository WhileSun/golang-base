package vo

type MdDocumentNameList struct {
	Id       int    `json:"id"`
	DocumentName  string `json:"document_name"`
	DocumentIdent string `json:"document_ident"`
	BookId        int    `json:"book_id"`
	ParentId      int    `json:"parent_id"`
	OrderSort     int    `json:"order_sort"`
}
