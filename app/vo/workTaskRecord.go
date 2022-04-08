package vo

import "github.com/whilesun/go-admin/app/po"

type WorkTaskRecordList struct {
	po.WorkTaskRecord
	ProjectName string `json:"project_name"`
}
