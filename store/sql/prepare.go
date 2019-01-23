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
	v = reflect.Indirect(v)
	return unmarshalValue(v, columns)
}

func unmarshalValue(v reflect.Value, columns []string) (dest []interface{}, err error) {
	switch v.Kind() {
	case reflect.Bool, reflect.String,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		dest = append(dest, v.Addr().Interface())
		return
	case reflect.Struct:
		if v.Type().Name() == "Time" {
			dest = append(dest, v.Addr().Interface())
			return
		}
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

func getAllFieldsEx(itf interface{}) (fields map[string]reflect.Kind) {
	t := reflect.TypeOf(itf)
	v := reflect.ValueOf(itf)

	if t.Kind() == reflect.Ptr {
		v = v.Elem()
		t = v.Type()
	}

	if t.Kind() != reflect.Struct {
		return
	}

	fields = make(map[string]reflect.Kind)
	for i := 0; i < t.NumField(); i++ {
		key := getFieldName(t.Field(i))
		if len(key) == 0 {
			continue
		}
		fields[key] = v.Field(i).Kind()
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

func inFields(x string, fields map[string]reflect.Kind) (bool, reflect.Kind) {
	for k, v := range fields {
		if k == x {
			return true, v
		}
	}
	return false, 0
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
		case reflect.Struct:
			if value.Type().Name() == "Time" {
				ret = !reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
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

	var keyList []string
	for i := 0; i < t.NumField(); i++ {
		key := getFieldName(t.Field(i))
		if len(key) == 0 {
			continue
		}
		value := v.Field(i)
		if value.IsValid() {
			keyList = append(keyList, key)
			args = append(args, value.Interface())
		}
	}
	query = joinKlist(keyList, sep)
	return
}

func joinKlist(keyList []string, sep string) string {
	l := len(keyList)
	var buff bytes.Buffer
	for n, key := range keyList {
		buff.WriteString(key)
		buff.WriteString("=?")
		if n < l-1 {
			buff.WriteString(sep)
		}
	}
	return buff.String()
}
