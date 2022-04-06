package gtools

import (
	"encoding/json"
	"reflect"
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