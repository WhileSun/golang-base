package gtools

import (
	"encoding/json"
	"fmt"
	"github.com/whilesun/go-admin/pkg/gcrypto"
	"math/rand"
	"reflect"
	"sort"
	"time"
	"unicode"
)


//实现三元表达式的功能
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	} else {
		return falseVal
	}
}

//json编码加转string
func JsonEncoode(res interface{}) string{
	jsonString, _ := json.Marshal(res)
	return string(jsonString)
}

//InArray 判断某一个值是否含在切片之中
func InArray(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(array)

			for i := 0; i < s.Len(); i++ {
				if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
					index = i
					exists = true
					return
				}
			}
	}

	return
}

//StructToMap struct转化为map
func StructToMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

//ArrStrUnique 一维string数组去重
func ArrStrUnique(arr []string)  (newArr []string){
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

func StringDefault(str string, defaultStr string) string{
	if str == ""{
		return defaultStr
	}else{
		return str
	}
}

func VerifyPasswdV4(s string) bool {
	var hasNumber, hasUpperCase, hasLowercase, hasSpecial bool
	for _, c := range s {
		switch {
		case unicode.IsNumber(c):
			hasNumber = true
		case unicode.IsUpper(c):
			hasUpperCase = true
		case unicode.IsLower(c):
			hasLowercase = true
		case c == '#' || c == '|':
			return false
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			hasSpecial = true
		}
	}
	return hasNumber && hasUpperCase && hasLowercase && hasSpecial
}

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
	sort.Strings(compArr)
	for _,tVal := range targetArr{
		index := sort.SearchStrings(compArr, tVal)
		if !(index < len(compArr) && compArr[index] == tVal) {
			notArr = append(notArr, tVal)
		}
	}
	return notArr
}

func StrArrayEquals(a []string , b []string, sortBool bool) bool{
	if len(a) != len(b){
		return false
	}
	if sortBool{
		sort.Strings(a)
		sort.Strings(b)
	}
	for i,val := range a{
		if val != b[i]{
			return false
		}
	}
	return true
}

func StrRand(header string) string{
	return header + gcrypto.Md5Encode16(fmt.Sprintf("%d-%d",time.Now().UnixNano(),rand.Intn(99999)))
}