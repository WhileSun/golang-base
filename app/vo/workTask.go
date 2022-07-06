package vo

import "github.com/whilesun/go-admin/app/po"

type WorkTaskList struct {
	po.WorkTask
	ProjectName string `json:"project_name"`
}
