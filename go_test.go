package main

import (
	"fmt"
	"sort"
	"testing"
)


func InStrArray(target string, strArray []string) bool{
	sort.Strings(strArray)
	index := sort.SearchStrings(strArray, target)
	if index < len(strArray) && strArray[index] == target {
		return true
	}
	return false
}

func InArrayNotExist(targetArr []string, compArr []string) []string{
	notArr:= make([]string,0)
	for _,tVal := range targetArr{
		if !InStrArray(tVal,compArr) {
			notArr = append(notArr, tVal)
		}
	}
	return notArr
}

func TestToken(t *testing.T) {
	fmt.Printf("%+v",InArrayNotExist([]string{"a","b"},[]string{"a","c"}))
}
