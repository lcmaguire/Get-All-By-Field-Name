package finder

import (
	"reflect"
)

func Find(iface interface{}, fieldName string) []string {
	out := make([]string, 0)
	rv := reflect.ValueOf(iface)

	if rv.Kind() == reflect.Slice {
		for i := 0; i < rv.Len(); i++ {
			out = append(out, find(rv.Index(i), fieldName))
		}
		return out
	}

	if rv.Kind() == reflect.Map {
		iter := rv.MapRange()
		for iter.Next() {
			out = append(out, find(iter.Value(), fieldName))
		}
		return out
	}

	return []string{find(rv, fieldName)}
}

func find(val reflect.Value, fieldName string) string {
	if val.Kind() == reflect.Pointer {
		return val.Elem().FieldByName(fieldName).String()
	}

	return val.FieldByName(fieldName).String()
}
