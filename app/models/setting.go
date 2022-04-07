package models

import "github.com/whilesun/go-admin/app/po"

type SSetting struct {
	po.SSetting
}

func NewSetting() *SSetting {
	return &SSetting{}
}

