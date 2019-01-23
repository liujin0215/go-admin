package crud

// RetrieveData 查询返回的data结构
type RetrieveData struct {
	Count  uint        `json:"count"`
	Record interface{} `json:"record"`
}
