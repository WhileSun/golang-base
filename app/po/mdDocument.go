package po

type MdDocument struct {
	BaseField
	DocumentName  string `json:"document_name"`
	DocumentIdent string `json:"document_ident"`
	BookId        int    `json:"book_id"`
	ParentId      int    `json:"parent_id"`
	OrderSort     int    `json:"order_sort"`
	MdText        string `json:"md_text"`
	HtmlText      string `json:"html_text"`
}
