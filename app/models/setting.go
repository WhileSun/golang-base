package models

import (
	po2 "github.com/whilesun/go-admin/app/types/po"
)

type SSetting struct {
	po2.SSetting
}

func NewSetting() *SSetting {
	return &SSetting{}
}
