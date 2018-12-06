package sql

import (
	"reflect"
)

func unmarshalStruct(itf interface{}, columns []string) (dest []interface{}, err error) {
	v := reflect.ValueOf(itf)
	if v.Kind() != reflect.Ptr {
		err = ErrNotPtr
		return
	}
	v = reflect.Indirect(reflect.ValueOf(itf))
	return unmarshalValue(v, columns)
}

func unmarshalValue(v reflect.Value, columns []string) (dest []interface{}, err error) {
	t := v.Type()
	for _, field := range columns {
		for i := 0; i < t.NumField(); i++ {
			ti := getFieldName(t.Field(i))
			if field != ti {
				continue
			}
			vi := v.Field(i)
			if vi.CanAddr() {
				dest = append(dest, vi.Addr().Interface())
			} else {
				err = ErrNotPtr
				return
			}
		}
	}
	return
}
