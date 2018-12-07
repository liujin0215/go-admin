package sql

import (
	"bytes"
	"reflect"
	"strings"
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

func prepareSelect(itf interface{}) string {
	t := reflect.TypeOf(itf)
	v := reflect.ValueOf(itf)

	if t.Kind() == reflect.String {
		return itf.(string)
	}

	if t.Kind() == reflect.Ptr {
		v = v.Elem()
		t = v.Type()
	}

	if t.Kind() != reflect.Struct {
		return ""
	}

	var buff bytes.Buffer
	for i := 0; i < t.NumField(); i++ {
		key := getFieldName(t.Field(i))
		if len(key) == 0 {
			continue
		}
		buff.WriteString(key)
		if i < t.NumField()-1 {
			buff.WriteString(",")
		}
	}
	return buff.String()
}

func getAllFields(itf interface{}) (fields []string) {
	t := reflect.TypeOf(itf)
	v := reflect.ValueOf(itf)

	if t.Kind() == reflect.Ptr {
		v = v.Elem()
		t = v.Type()
	}

	if t.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < t.NumField(); i++ {
		key := getFieldName(t.Field(i))
		if len(key) == 0 {
			continue
		}
		fields = append(fields, key)
	}
	return
}

func getFieldName(field reflect.StructField) string {
	tag := field.Tag.Get(tagName)
	if len(tag) == 0 {
		return ""
	}
	options := strings.Split(tag, ",")
	return strings.TrimSpace(options[0])
}

func prepareWhere(where interface{}, args []interface{}) (string, []interface{}) {
	if query, ok := where.(string); ok {
		return query, args
	}

	if model, ok := where.(MapModel); ok {
		return prepareMapQuery(model, " and ")
	}

	return "", nil
}

func parseMap(mapModel map[string]interface{}) (kList []string, vList []interface{}) {
	for key, value := range mapModel {
		kList = append(kList, key)
		vList = append(vList, value)
	}
	return
}

func instrList(x string, strList []string) bool {
	for _, i := range strList {
		if i == x {
			return true
		}
	}
	return false
}

func prepareMapQuery(model map[string]interface{}, sep string) (query string, args []interface{}) {
	var kList []string
	kList, args = parseMap(model)
	query = joinKlist(kList, sep)
	return
}

func joinKlist(kList []string, sep string) string {
	l := len(kList)
	var buff bytes.Buffer
	for n, key := range kList {
		buff.WriteString(key)
		buff.WriteString(" = ?")
		if n < l-1 {
			buff.WriteString(sep)
		}
	}
	return buff.String()
}
