package utils

import (
	"fmt"
	"reflect"
	"strings"
)

type ColumnArgs struct {
	Query string
	Args  []interface{}
}

func ParseColumn(s interface{}, paramsIndex int, dest *ColumnArgs) ColumnArgs {
	val := reflect.ValueOf(s)
	num := val.NumField()

	for i := 0; i < num; i++ {
		f := val.Field(i)
		kind := f.Kind()
		if kind == reflect.Ptr {
			if !f.IsNil() {
				t := val.Type()
				sField, ok := t.Field(i).Tag.Lookup("db")
				if !ok {
					continue
				}
				_, ok = t.Field(i).Tag.Lookup("skip")
				if ok {
					continue
				}

				dest.Query += fmt.Sprintf(" %s = $%d,", sField, len(dest.Args)+paramsIndex)
				dest.Args = append(dest.Args, f.Elem().Interface())
			}
		} else {
			if !f.IsZero() {
				t := val.Type()
				sField, ok := t.Field(i).Tag.Lookup("db")
				if !ok {
					continue
				}
				_, ok = t.Field(i).Tag.Lookup("skip")
				if ok {
					continue
				}

				dest.Query += fmt.Sprintf(" %s = $%d,", sField, len(dest.Args)+paramsIndex)
				dest.Args = append(dest.Args, f.Interface())
			}
		}
	}

	dest.Query = strings.ReplaceAll(dest.Query, "\t", "")
	dest.Query = strings.Trim(dest.Query, ",")

	return *dest
}
