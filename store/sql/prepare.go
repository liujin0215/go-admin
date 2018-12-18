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
	switch itf.(type) {
	case string:
		return itf.(string)
	default:
		fields := getAllFields(itf)
		return strings.Join(fields, ",")
	}
	return ""
}

func prepareWhere(where interface{}, args []interface{}) (string, []interface{}) {
	switch where.(type) {
	case string:
		return where.(string), args
	case MapModel:
		if len(args) == 0 {
			return "", args
		}
		model, ok := args[1].(Model)
		if !ok {
			return "", []interface{}{}
		}
		keyList, valueList := where.(MapModel).Parse(model)
		return joinKlist(keyList, " and "), valueList
	default:
		keyList, valueList := parseStruct(where)
		return joinKlist(keyList, " and "), valueList
	}
	return "", []interface{}{}
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

func instrList(x string, strList []string) bool {
	for _, i := range strList {
		if i == x {
			return true
		}
	}
	return false
}

func parseStruct(itf interface{}) (keyList []string, valueList []interface{}) {
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
		value := v.Field(i)
		if !value.IsValid() {
			continue
		}
		var ret bool
		switch value.Kind() {
		case reflect.String:
			if len(value.String()) > 0 {
				ret = true
			}
			break
		case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int8:
			if value.Int() > 0 {
				ret = true
			}
			break
		case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint8:
			if value.Uint() > 0 {
				ret = true
			}
			break
		default:
			ret = true
		}
		if ret {
			keyList = append(keyList, key)
			valueList = append(valueList, value.Interface())
		}
	}
	return
}

func prepareStructQuery(model interface{}, sep string) (query string, args []interface{}) {
	t := reflect.TypeOf(model)
	v := reflect.ValueOf(model)

	if t.Kind() == reflect.Ptr {
		v = v.Elem()
		t = v.Type()
	}

	if t.Kind() != reflect.Struct {
		return
	}

	var kList []string
	for i := 0; i < t.NumField(); i++ {
		key := getFieldName(t.Field(i))
		if len(key) == 0 {
			continue
		}
		value := v.Field(i)
		if value.IsValid() {
			kList = append(kList, key)
			args = append(args, value.Interface())
		}
	}
	query = joinKlist(kList, sep)
	return
}

func joinKlist(kList []string, sep string) string {
	l := len(kList)
	var buff bytes.Buffer
	for n, key := range kList {
		buff.WriteString(key)
		buff.WriteString("=?")
		if n < l-1 {
			buff.WriteString(sep)
		}
	}
	return buff.String()
}
