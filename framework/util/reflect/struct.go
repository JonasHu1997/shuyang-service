package reflect

import (
	"gin-service/framework/util"
	"reflect"
	"strings"
)

type RefStruct struct {
	dat      interface{}
	refValue reflect.Value
	fieldNum int
	tagIdx   map[string]int
}

func NewRefStruct(dat interface{}, tagName string) *RefStruct {
	refValue := reflect.ValueOf(dat).Elem()
	if refValue.Kind() != reflect.Struct {
		panic("require struct")
	}
	fieldNum := refValue.NumField()
	if fieldNum == 0 {
		panic("fields num is 0")
	}
	tagIdx := make(map[string]int, fieldNum)
	for i := 0; i <= fieldNum; i++ {
		tag := ParseTag(tagName, refValue.Type().Field(i))
		if tag != "" {
			tagIdx[tag] = i
		}
	}
	return &RefStruct{dat, refValue, fieldNum, tagIdx}
}

func (r *RefStruct) SetField(field, value string) {
	idx := 0
	ok := true
	if idx, ok = r.tagIdx[field]; !ok {
		return
	}
	if !r.refValue.Field(idx).CanSet() {
		return
	}
	t := r.refValue.Field(idx).Kind()
	switch t {
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		intVal, _ := util.ToInt(value)
		r.refValue.Field(idx).Set(reflect.ValueOf(intVal))
	case reflect.String:
		r.refValue.Field(idx).Set(reflect.ValueOf(value))
	}
}

func (r *RefStruct) TagPtr() (ret map[string]interface{}) {
	l := len(r.tagIdx)
	if l == 0 {
		return
	}
	ret = make(map[string]interface{}, l)
	for tag, i := range r.tagIdx {
		ptr := r.refValue.Field(i).Addr().Interface()
		ret[tag] = ptr
	}
	return
}

func ParseTag(tag string, refField reflect.StructField) string {
	field := refField.Tag.Get(tag)
	field = strings.TrimSpace(field)
	if field == "" {
		return ""
	}
	if field != "-" {
		t := strings.Split(field, ",")
		return t[0]
	}
	return ""
}
