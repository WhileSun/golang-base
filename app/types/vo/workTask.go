package vo

import (
	po2 "github.com/whilesun/go-admin/app/types/po"
)

type WorkTaskList struct {
	po2.WorkTask
	ProjectName string `json:"project_name"`
}
