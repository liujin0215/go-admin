package sql

// 常量
const (
	//tagName 用于解析对应数据库的字段名
	tagName = "model"
)

// 数据库操作方法的枚举
const (
	InsertMethod = iota + 1
	UpdateMethod
	DeleteMethod
	SelectMethod
	SelectOneMethod
)
