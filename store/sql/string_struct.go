package sql

type StringStruct struct {
	stmt string
	args []interface{}
}

func NewStringStruct(s string, args ...interface{}) *StringStruct {
	return &StringStruct{s, args}
}

func (ss *StringStruct) Stmt() string {
	return ss.stmt
}

func (ss *StringStruct) Args() []interface{} {
	return ss.args
}
