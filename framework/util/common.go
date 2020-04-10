package util

import (
	"fmt"
	"strconv"
)

// ToString 任意类型装换为string
func ToString(arg interface{}) (ret string, err error) {
	switch arg.(type) {
	case int:
		ret = strconv.Itoa(arg.(int))
	case int64:
		ret = strconv.FormatInt(arg.(int64), 10)
	case []byte:
		ret = string(arg.([]byte))
	case string:
		ret = arg.(string)
	case float64:
		ret = strconv.FormatFloat(arg.(float64), 'f', -2, 64)
	case float32:
		ret = strconv.FormatFloat(arg.(float64), 'f', -2, 64)
	default:
		err = fmt.Errorf("Unsupport Type")
	}
	return
}

// ToInt 任意类型转换为int
func ToInt(value interface{}) (ret int, err error) {
	switch value.(type) {
	case int:
		ret = value.(int)
	case int8:
		ret = int(value.(int8))
	case int16:
		ret = int(value.(int16))
	case int32:
		ret = int(value.(int32))
	case int64:
		ret = int(value.(int64))
	case uint:
		ret = int(value.(uint))
	case uint8:
		ret = int(value.(uint8))
	case uint16:
		ret = int(value.(uint16))
	case uint32:
		ret = int(value.(uint32))
	case uint64:
		ret = int(value.(uint64))
	case string:
		ret, err = strconv.Atoi(value.(string))
	default:
		err = fmt.Errorf("Unsupported Type")
	}
	return
}
