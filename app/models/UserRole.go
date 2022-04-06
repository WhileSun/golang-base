package models

import "github.com/whilesun/go-admin/app/po"

type RUserRole struct {
	po.RUserRole
}

func NewUserRole() *RUserRole {
	return &RUserRole{}
}
