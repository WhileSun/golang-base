package po

// WorkProject 工作任务项目
type WorkProject struct {
	BaseField
	ProjectName string `json:"project_name"`
	CreaterId   int    `json:"creater_id"`
	Remark      string `json:"remark"`
}
