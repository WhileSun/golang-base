package dto

type AddWorkTaskProject struct {
	ProjectName string `form:"project_name"  binding:"required" label:"项目名称"`
	Remark      string `form:"remark"  binding:"" label:"项目备注"`
}

type UpdateWorkTaskProject struct {
	Id int `form:"id"  binding:"required" label:"ID"`
	AddWorkTaskProject
}
