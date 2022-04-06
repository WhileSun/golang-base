package dto

type AddTaskRecord struct {
	TaskProjectId int    `form:"task_project_id"  binding:"required" label:"项目ID"`
	LaunchTimeStr string `form:"launch_time" binding:"required,datetime=2006-01-02 15:04:05" label:"发起时间"`
	StartTimeStr  string `form:"start_time" binding:"omitempty,datetime=2006-01-02 15:04:05" label:"开始时间"`
	EndTimeStr    string `form:"end_time"  binding:"omitempty,datetime=2006-01-02 15:04:05"  label:"结束时间"`
	TaskType      int16  `form:"task_type" binding:"required" label:"任务类型"`
	TaskLevel     int16  `form:"task_level" binding:"required" label:"优先级"`
	PerformStatus int16  `form:"perform_status" binding:"required" label:"任务状态"`
	Title         string `form:"title" binding:"required" label:"标题"`
	Content       string `form:"content" binding:"required" label:"任务内容"`
	Remark        string `form:"remark" binding:"" label:"备注"`
}

type UpdateTaskRecord struct {
	Id int `form:"id"  binding:"required" label:"ID"`
	AddTaskRecord
}