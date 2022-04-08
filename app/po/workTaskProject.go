package po

// WorkTaskProject 工作任务项目
type WorkTaskProject struct {
	BaseField
	ProjectName string `json:"project_name"`
	CreaterId   int    `json:"creater_id"`
	Remark      string `json:"remark"`
}
