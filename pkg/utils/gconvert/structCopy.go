package gconvert

import (
	"errors"
	"fmt"
	"reflect"
)

func StructCopy(src, dst interface{},otherTypes ...string) error {
	srcV, err := srcFilter(src)
	//fmt.Printf("srcV--%+v \n", srcV)
	if err != nil {
		return err
	}
	dstV, err := dstFilter(dst)
	//fmt.Printf("dstV--%+v \n", dstV)
	if err != nil {
		return err
	}
	srcKeys := make(map[string]bool)
	types := filterType(otherTypes)
	getSrcKeys(srcV, srcKeys, types)
	setDstVals(dstV, srcV, srcKeys, types)
	return nil
}

func filterType(otherTypes []string) []string {
	types := []string{"dbtypes.JSONTime", "gorm.DeletedAt"}
	types = append(types, otherTypes...)
	return types
}

func inArray(types []string, nowType string) bool {
	result := false
	for _, val := range types {
		if val == nowType {
			result = true
			break
		}
	}
	return result
}

func getSrcKeys(srcV reflect.Value, srcKeys map[string]bool, types []string) {
	for i := 0; i < srcV.NumField(); i++ {
		if srcV.Type().Field(i).Type.Kind() == reflect.Struct {
			if inArray(types, fmt.Sprintf("%s", srcV.Field(i).Type())) {
				srcKeys[srcV.Type().Field(i).Name] = true
				continue
			}
			subStruct := srcV.FieldByName(srcV.Type().Field(i).Name)
			getSrcKeys(subStruct, srcKeys, types)
		} else {
			srcKeys[srcV.Type().Field(i).Name] = true
		}
	}
}

func setDstVals(dstV reflect.Value, srcV reflect.Value, srcKeys map[string]bool, types []string) {
	for i := 0; i < dstV.NumField(); i++ {
		if dstV.Type().Field(i).Type.Kind() == reflect.Struct {
			//某些子struct不需要解析，直接替换
			if !inArray(types, fmt.Sprintf("%s", dstV.Field(i).Type())) {
				subStruct := dstV.FieldByName(dstV.Type().Field(i).Name)
				setDstVals(subStruct, srcV, srcKeys, types)
			}
		}
		fName := dstV.Type().Field(i).Name
		if _, ok := srcKeys[fName]; ok {
			v := srcV.FieldByName(dstV.Type().Field(i).Name)
			// 指针数据copy
			if v.Type().Kind() == reflect.Ptr {
				if v.IsNil() {
					continue
				}
				//判断复制的dstV是否也是指针类别
				v1 := dstV.FieldByName(dstV.Type().Field(i).Name)
				if v1.Type().Kind() != reflect.Ptr {
					v = v.Elem()
				}
			}
			if v.CanInterface() {
				dstV.Field(i).Set(v)
			}
		}
	}
}

func srcFilter(src interface{}) (reflect.Value, error) {
	v := reflect.ValueOf(src)
	if v.Type().Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return reflect.Zero(v.Type()), errors.New("src type error: not a struct or a pointer to struct")
	}
	return v, nil
}

func dstFilter(src interface{}) (reflect.Value, error) {
	v := reflect.ValueOf(src)
	if v.Type().Kind() == reflect.Ptr {
		v = v.Elem()
		//return reflect.Zero(v.Type()), errors.New("src type error: not a pointer to struct")
	}
	if v.Kind() != reflect.Struct {
		return reflect.Zero(v.Type()), errors.New("src type error: not point to struct")
	}
	return v, nil
}
