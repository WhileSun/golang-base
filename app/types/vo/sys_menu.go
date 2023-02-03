package vo

// SysMenuModel s_menu 菜单栏目
type SysMenuModel struct {
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

type MenuFieldList struct {
	Id       int    `json:"id"`
	MenuName string `json:"menu_name"`
	ParentId int    `json:"parent_id"`
	Status   bool   `json:"status"`
	MenuType int    `json:"menu_type"`
}
