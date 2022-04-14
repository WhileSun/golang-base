package po

/**
SMenu 菜单表
menu_type 1 菜单 2 节点
*/

type SMenu struct {
	BaseField
	MenuName  string `json:"menu_name"`
	MenuType  int16  `json:"menu_type"`
	Url       string `json:"url"`
	Icon      string `json:"icon"`
	ParentId  int    `json:"parent_id"`
	Sort      int    `json:"sort"`
	DataPerms string `json:"data_perms"`
	PagePerms string `json:"page_perms"`
	Status    bool   `json:"status" gorm:"default:true;"`
	Show      bool   `json:"show" gorm:"default:true;"`
	IsSys     bool   `json:"is_sys" gorm:"default:false;"`
}
