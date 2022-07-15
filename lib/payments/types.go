package payments

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// change string to type of Type
func ChangeType(s string, Type reflect.Type) (reflect.Value, error) {
	String := reflect.TypeOf("")
	Time1 := reflect.TypeOf(time.Now())
	Float64 := reflect.TypeOf(float64(0))
	Float32 := reflect.TypeOf(float32(0))
	Int := reflect.TypeOf(int(0))
	Int32 := reflect.TypeOf(int32(0))
	Int64 := reflect.TypeOf(int64(0))
	Time2 := reflect.TypeOf(Time{})
	Bool := reflect.TypeOf(false)

	switch Type {
	case String:
		return reflect.ValueOf(s), nil
	case Time1:
		value, _ := time.Parse(time.RFC3339, s)
		return reflect.ValueOf(value), nil
	case Float64:
		value, _ := strconv.ParseFloat(s, 64)
		return reflect.ValueOf(value), nil
	case Float32:
		value, _ := strconv.ParseFloat(s, 32)
		return reflect.ValueOf(float32(value)), nil
	case Int:
		value, _ := strconv.Atoi(s)
		return reflect.ValueOf(value), nil
	case Int32:
		value, _ := strconv.Atoi(s)
		return reflect.ValueOf(int32(value)), nil
	case Int64:
		value, _ := strconv.Atoi(s)
		return reflect.ValueOf(int64(value)), nil
	case Bool:
		value, _ := strconv.ParseBool(s)
		return reflect.ValueOf(value), nil
	case Time2:
		var Time Time
		Time.UnmarshalJSON([]byte(s))
		return reflect.ValueOf(Time), nil
	default:
		return reflect.ValueOf(s), errors.New("type not found")
	}
}


func GetFieldIndex(Struct reflect.Value, tag string) int {
	for index:=0; index<Struct.Type().NumField(); index++ {
		name := Struct.Type().Field(index).Tag.Get("json")
		name = strings.ReplaceAll(name, ",omitempty", "")
		if name == tag {
			return index
		}
	}
	return -1
}