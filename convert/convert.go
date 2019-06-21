package convert

import (
	"github.com/flxxyz/util"
	"reflect"
	"strconv"
)

var structToMapHayStack = []interface{}{
	"string",
	"int",
	"int8",
	"int16",
	"int32",
	"int64",
	"float32",
	"float64",
	"uint",
	"uint8",
	"uint16",
	"uint32",
	"uint64",
	"bool",
	"byte",
}

// struct转换map
func StructToMap(data interface{}) map[string]interface{} {
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)

	mapConvert := make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		temp := v.Field(i).Interface()
		if !util.InArray(v.Field(i).Type().Name(), structToMapHayStack) {
			temp = StructToMap(temp)
		}
		mapConvert[t.Field(i).Name] = temp
	}
	return mapConvert
}

// struct转换json
func StructToJson() {

}

// int转换string
func IntToStr(v int) string {
	return strconv.Itoa(v)
}

// int转换int64
func IntTo64(v int) int64 {
	return int64(v)
}

// string转换int
func StrToInt(v string) int {
	i, _ := strconv.Atoi(v)
	return i
}

// string转换int64
func StrTo64(v string) (i int64, err error) {
	i, err = strconv.ParseInt(v, 10, 64)
	return
}

// string转换bool
func StrToBool(v string) (b bool) {
	b, _ = strconv.ParseBool(v)
	return
}

// string转换byte
func StrToByte(v string) uint8 {
	i, _ := strconv.ParseUint(v, 10, 8)
	return uint8(i)
}

// bool转换string
func BoolToStr(v bool) string {
	return strconv.FormatBool(v)
}
