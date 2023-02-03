package po

import (
	"github.com/whilesun/go-admin/pkg/utils/gtime"
)

type WorkTask struct {
	BaseField
	ProjectId     int            `json:"project_id"`
	LaunchTime    gtime.DateTime `json:"launch_time"`
	StartTime     gtime.DateTime `json:"start_time"`
	EndTime       gtime.DateTime `json:"end_time"`
	TaskType      int16          `json:"task_type"`
	TaskLevel     int16          `json:"task_level"`
	PerformStatus int16          `json:"perform_status"`
	Title         string         `json:"title"`
	Content       string         `json:"content"`
	Remark        string         `json:"remark"`
}
