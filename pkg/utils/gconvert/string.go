package gconvert

import "strconv"

func StrToInt(str string) int {
	resp, _ := strconv.Atoi(str)
	return resp
}

func StrToBool(str string) bool {
	resp, _ := strconv.ParseBool(str)
	return resp
}
