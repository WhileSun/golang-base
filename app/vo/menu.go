package vo

type MenuFieldList struct {
	Id       int    `json:"id"`
	MenuName string `json:"menu_name"`
	ParentId int	`json:"parent_id"`
	Status   bool	`json:"status"`
	MenuType int    `json:"menu_type"`
}
