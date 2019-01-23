package sql

import (
	"database/sql"
	"reflect"
)

// 定义结构体
type (
	// DB 数据库操作类
	DB struct {
		*sql.DB
	}

	// DBConfig 数据库配置项
	DBConfig struct {
		Type string
		Addr string
	}
)

// NewDB 通过数据库配置项生成一个新的数据库操作类
func NewDB(c *DBConfig) (db *DB, err error) {
	db = new(DB)
	db.DB, err = sql.Open(c.Type, c.Addr)
	return
}

// Exec 对用数据库操作的exec，用于insert,update,delete
func (db *DB) Exec(method int, es ExecStruct) (sql.Result, error) {
	stmt := es.Prepare(method)
	return db.DB.Exec(stmt.Stmt(), stmt.Args()...)
}

// Retrieve 数据库批量查询操作的方法
func (db *DB) Retrieve(ms *ModelStruct) (res interface{}, err error) {
	stmt := ms.Prepare(SelectMethod)
	res = reflect.New(reflect.SliceOf(reflect.ValueOf(ms.Model).Type())).Interface()
	err = db.Raw(stmt.Stmt(), stmt.Args()...).Find(res).Err()
	return
}

// RetrieveOne 数据库单条数据查询的操作方法
func (db *DB) RetrieveOne(m Model) (err error) {
	return db.Tb(m.TbName()).Select(m).Where(m).FindOne(m).Err()
}

// RetrieveOneEx 通过ModelStruct进行数据库单条数据查询的操作方法
func (db *DB) RetrieveOneEx(ms *ModelStruct) (res interface{}, err error) {
	stmt := ms.Prepare(SelectMethod)
	//res = ms.Model
	res = reflect.New(reflect.Indirect(reflect.ValueOf(ms.Model)).Type()).Interface()
	err = db.Raw(stmt.Stmt(), stmt.Args()...).FindOne(res).Err()
	return
}

// Select 通过select方法生成query类
func (db *DB) Select(itf interface{}) *Query {
	q := &Query{selectQuery: prepareSelect(itf), db: db.DB}
	if len(q.selectQuery) == 0 {
		q.err = ErrQueryEmpty
	}
	return q
}

// Tb 通过Tb方法生成query类
func (db *DB) Tb(tbName string) *Query {
	return &Query{tb: tbName, db: db.DB}
}

// Raw 通过Raw方法生成query类
func (db *DB) Raw(query string, args ...interface{}) *Query {
	return &Query{query: query, db: db.DB, args: args}
}
