package sql

type (
	MapModel map[string]interface{}
)

func (mm MapModel) GetFloat(key string) (val float64, ok bool) {
	var itf interface{}
	itf, ok = mm[key]
	if !ok {
		return
	}

	val, ok = itf.(float64)
	return
}

func (mm MapModel) GetString(key string) (val string) {
	itf, ok := mm[key]
	if !ok {
		return
	}

	val, _ = itf.(string)
	return
}

func (mm MapModel) GetUint(key string) uint {
	v, ok := mm.GetFloat(key)
	if !ok || v <= 0 {
		return 0
	}

	return uint(v)
}

func (mm MapModel) PK() uint {
	return mm.GetUint("id")
}

func (mm MapModel) ClearPK() {
	delete(mm, "id")
}

func (mm MapModel) OrderBy() string {
	return mm.GetString("order_by")
}

func (mm MapModel) Limit() uint {
	pagesize := mm.GetUint("pagesize")
	if pagesize > 20 {
		return 20
	}
	return pagesize
}

func (mm MapModel) Offset() uint {
	pagenum := mm.GetUint("pagenum")
	if pagenum == 0 {
		pagenum = 1
	}
	limit := mm.Limit()
	return limit * (pagenum - 1)
}

func (mm MapModel) Parse(itf interface{}) (keyList []string, valueList []interface{}) {
	fields := getAllFields(itf)
	for key, value := range mm {
		if instrList(key, fields) {
			keyList = append(keyList, key)
			valueList = append(valueList, value)
		}
	}
	return
}
