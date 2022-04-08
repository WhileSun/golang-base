package gconvert

import (
	"strconv"
	"time"
)
const (
	YYYYMMDDHHMISS = "2006-01-02 15:04:05" //常规类型
)

func StrToInt(str string) int {
	resp, _ := strconv.Atoi(str)
	return resp
}

func StrToBool(str string) bool {
	resp, _ := strconv.ParseBool(str)
	return resp
}

func StrToDatetime(str string) time.Time{
	st, _  := time.Parse(YYYYMMDDHHMISS, str)
	return st
}