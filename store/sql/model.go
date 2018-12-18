package sql

type Model interface {
	TbName() string
	PK() uint
}

type StmtStruct interface {
	Stmt() string
	Args() []interface{}
}

type ExecStruct interface {
	Prepare(int) StmtStruct
}
