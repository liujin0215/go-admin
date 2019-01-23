package response

// 错误的枚举
const (
	ErrCreate = iota + 1
	ErrUpdate
	ErrDelete
	ErrRetrieve
	ErrParseBody
	ErrEmptyPK
)

// 错误的返回内容map
var errMap = map[int]string{
	ErrParseBody: "解析body失败",
	ErrCreate:    "创建失败",
	ErrRetrieve:  "查询失败",
	ErrEmptyPK:   "主键为空",
	ErrUpdate:    "更新失败",
	ErrDelete:    "删除失败",
}
