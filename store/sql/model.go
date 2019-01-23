// Package sql 各种模型的抽象接口
package sql

// Model 基础数据库模型的接口
type Model interface {
	TbName() string
	PK() uint
}

// StmtStruct 数据库语句的接口
type StmtStruct interface {
	Stmt() string
	Args() []interface{}
}

// ExecStruct 数据库准备的接口，通过prepare方法生成数据库语句的接口
type ExecStruct interface {
	Prepare(int) StmtStruct
}
