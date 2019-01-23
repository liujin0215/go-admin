// Package sql StringStruct类
package sql

// StringStruct StringStruct类的结构
// StmtStruct的令一种实现
type StringStruct struct {
	stmt string
	args []interface{}
}

// NewStringStruct 生成StringStruct类
func NewStringStruct(s string, args ...interface{}) *StringStruct {
	return &StringStruct{s, args}
}

// Stmt 实现StmtStruct
func (ss *StringStruct) Stmt() string {
	return ss.stmt
}

// Args 实现StmtStruct
func (ss *StringStruct) Args() []interface{} {
	return ss.args
}
