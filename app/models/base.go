package models

import (
	"github.com/whilesun/go-admin/pkg/gsys"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
)

func init() {
	db = gsys.Db
}


