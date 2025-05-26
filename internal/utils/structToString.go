package utils

import (
	"fmt"
	"reflect"
	"strings"
)

func StructToString(s interface{}) string {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	var items []string
	for i := 0; i < t.NumField(); i++ {
		fieldName := t.Field(i).Name
		fieldValue := v.Field(i)
		items = append(items, fmt.Sprintf("%s: %v", fieldName, fieldValue))
	}

	return strings.Join(items, ", ")
}
