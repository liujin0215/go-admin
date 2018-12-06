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

func (mm MapModel) PK() uint {
	floatID, ok := mm.GetFloat("id")
	if !ok || floatID <= 0 {
		return 0
	}

	return uint(floatID)
}
