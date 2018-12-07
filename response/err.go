package response

const (
	ErrParseBody = iota + 1
	ErrCreate
	ErrRetrieve
	ErrEmptyPK
	ErrUpdate
	ErrDelete
)

var errMap = map[int]string{
	ErrParseBody: "解析body失败",
	ErrCreate:    "创建失败",
	ErrRetrieve:  "查询失败",
	ErrEmptyPK:   "主键为空",
	ErrUpdate:    "更新失败",
	ErrDelete:    "删除失败",
}
