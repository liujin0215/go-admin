package sql

import (
	"database/sql"
)

type (
	DB struct {
		*sql.DB
	}

	DBConfig struct {
		Type string
		Addr string
	}
)

func NewDB(c *DBConfig) (db *DB, err error) {
	db = new(DB)
	db.DB, err = sql.Open(c.Type, c.Addr)
	return
}

func (db *DB) Exec(method int, es ExecStruct) (sql.Result, error) {
	stmt := es.Prepare(method)
	return db.DB.Exec(stmt.Stmt(), stmt.Args()...)
}

func (db *DB) Select(itf interface{}) *Query {
	q := &Query{selectQuery: prepareSelect(itf), db: db.DB}
	if len(q.selectQuery) == 0 {
		q.err = ErrQueryEmpty
	}
	return q
}

func (db *DB) Tb(tbName string) *Query {
	return &Query{tb: tbName, db: db.DB}
}

func (db *DB) Raw(query string) *Query {
	return &Query{query: query, db: db.DB}
}
