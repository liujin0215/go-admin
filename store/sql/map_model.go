//Package sql MapModel类
package sql

import "reflect"

type (
	// MapModel 本质是map
	MapModel map[string]interface{}
)

// GetFloat 解析浮点型字段
func (mm MapModel) GetFloat(key string) (val float64, ok bool) {
	var itf interface{}
	itf, ok = mm[key]
	if !ok {
		return
	}

	val, ok = itf.(float64)
	return
}

// GetString 解析字符串型字段
func (mm MapModel) GetString(key string) (val string) {
	itf, ok := mm[key]
	if !ok {
		return
	}

	val, _ = itf.(string)
	return
}

// GetUint 解析无符号整型字段
func (mm MapModel) GetUint(key string) uint {
	v, ok := mm.GetFloat(key)
	if !ok || v <= 0 {
		return 0
	}

	return uint(v)
}

// GetInt 解析整型字段
func (mm MapModel) GetInt(key string) int64 {
	v, ok := mm.GetFloat(key)
	if !ok {
		return 0
	}

	return int64(v)
}

// PK 解析主键
func (mm MapModel) PK() uint {
	return mm.GetUint("id")
}

// ClearPK 清除主键
func (mm MapModel) ClearPK() {
	delete(mm, "id")
}

// OrderBy 解析排序字段
func (mm MapModel) OrderBy() string {
	return mm.GetString("order_by")
}

// Limit 解析limit
func (mm MapModel) Limit() uint {
	pagesize := mm.GetUint("pagesize")
	if pagesize > 20 {
		return 20
	}
	return pagesize
}

// Offset 解析Offset
func (mm MapModel) Offset() uint {
	pagenum := mm.GetUint("pagenum")
	if pagenum == 0 {
		pagenum = 1
	}
	limit := mm.Limit()
	return limit * (pagenum - 1)
}

// Parse 解析字段
func (mm MapModel) Parse(itf interface{}) (keyList []string, valueList []interface{}) {
	fields := getAllFieldsEx(itf)
	var in bool
	var kind reflect.Kind
	for key := range mm {
		in, kind = inFields(key, fields)
		if !in {
			continue

		}
		keyList = append(keyList, key)

		// 整型字段必须特殊处理，否则数据库操作在数据过大时会有错误
		switch kind {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			valueList = append(valueList, mm.GetInt(key))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			valueList = append(valueList, mm.GetUint(key))
		default:
			valueList = append(valueList, mm[key])
		}
	}
	return
}
