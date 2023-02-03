package models

import (
	gsys2 "github.com/whilesun/go-admin/gctx"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func init() {
	db = gsys2.Db
}
