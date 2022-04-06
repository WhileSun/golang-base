package po

import (
	"github.com/whilesun/go-admin/pkg/utils/gtime"
)

type BaseField struct {
	Id        int            `gorm:"primary_key"  json:"id"`
	CreatedAt gtime.DateTime `json:"created_at"`
	UpdatedAt gtime.DateTime `json:"updated_at"`
}