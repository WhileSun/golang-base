package po

type MdBook struct {
	BaseField
	BookName  string `json:"book_name"`
	BookIdent string `json:"book_ident"`
}
