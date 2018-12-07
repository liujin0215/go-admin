package sql

import (
	"database/sql"
	"fmt"
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

func (db *DB) Insert(tbName string, model MapModel) (id int64, err error) {
	var result sql.Result
	query, args := prepareMapQuery(model, ",")
	if len(args) == 0 {
		return 0, ErrQueryEmpty
	}
	query = fmt.Sprintf("insert into %s set %s;", tbName, query)
	result, err = db.DB.Exec(query, args...)
	return result.LastInsertId()
}

func (db *DB) Update(tbName string, model MapModel, where interface{}, wArgs ...interface{}) (affected int64, err error) {
	model.ClearPK()
	sQuery, args := prepareMapQuery(model, ",")
	if len(args) == 0 {
		return 0, ErrQueryEmpty
	}
	wQuery, wArgs := prepareWhere(where, wArgs)
	if len(wArgs) == 0 {
		return 0, ErrConditionEmpty
	}
	args = append(args, wArgs...)
	var result sql.Result
	query := fmt.Sprintf("update %s set %s where %s;", tbName, sQuery, wQuery)
	fmt.Println(query)
	result, err = db.DB.Exec(query, args...)
	if err != nil {
		return
	}
	return result.RowsAffected()
}

func (db *DB) Delete(tbName string, where interface{}, wArgs ...interface{}) (affected int64, err error) {
	wQuery, wArgs := prepareWhere(where, wArgs)
	if len(wArgs) == 0 {
		return 0, ErrConditionEmpty
	}
	var result sql.Result
	query := fmt.Sprintf("delete from %s where %s;", tbName, wQuery)
	result, err = db.DB.Exec(query, wArgs...)
	if err != nil {
		return
	}
	return result.RowsAffected()
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
