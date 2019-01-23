// Package sql Query类
package sql

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
)

type (
	// Query Query类的结构
	Query struct {
		db          *sql.DB
		tb          string
		selectQuery string
		where       string
		orderby     []string
		limit       uint
		offset      uint
		query       string
		args        []interface{}
		err         error
	}
)

// Stmt 返回数据库查询的语句，实现StmtStruct
func (q *Query) Stmt() string { return q.query }

// Args 返回数据库查询的参数，实现StmtStruct
func (q *Query) Args() []interface{} { return q.args }

// Err 返回错误
func (q *Query) Err() error {
	return q.err
}

// Prepare 准备数据库查询的语句，同时实现ExecStruct
func (q *Query) Prepare() *Query {
	if q.err != nil {
		return q
	}

	if len(q.query) > 0 {
		return q
	}

	if len(q.tb) == 0 || len(q.selectQuery) == 0 {
		q.err = ErrQueryEmpty
		return q
	}

	q.query = fmt.Sprintf("select %s from %s", q.selectQuery, q.tb)
	if len(q.where) > 0 {
		q.query = fmt.Sprintf("%s where %s", q.query, q.where)
	}

	if len(q.orderby) > 0 {
		q.query += fmt.Sprintf(" order by %s", strings.Join(q.orderby, ","))
	}

	if q.limit > 0 {
		q.query += " limit ?"
		q.args = append(q.args, q.limit)
	}

	if q.offset > 0 {
		q.query += " offset ?"
		q.args = append(q.args, q.offset)
	}

	q.query += ";"
	return q
}

// Tb 设置表名称
func (q *Query) Tb(tbName string) *Query {
	if q.err != nil {
		return q
	}
	q.tb = tbName
	return q
}

// Select 设置select语句
func (q *Query) Select(itf interface{}) *Query {
	if q.err != nil {
		return q
	}
	q.selectQuery = prepareSelect(itf)
	if len(q.selectQuery) == 0 {
		q.err = ErrQueryEmpty
	}
	return q
}

// Where 设置where语句
func (q *Query) Where(where interface{}, args ...interface{}) *Query {
	if q.err != nil {
		return q
	}
	q.where, q.args = prepareWhere(where, args)
	return q
}

// OrderBy 设置order by语句
func (q *Query) OrderBy(args ...string) *Query {
	if q.err != nil {
		return q
	}
	q.orderby = append(q.orderby, args...)
	return q
}

// Limit 设置limit语句
func (q *Query) Limit(n uint) *Query {
	if q.err != nil {
		return q
	}
	q.limit = n
	return q
}

// Offset 设置offset 语句
func (q *Query) Offset(n uint) *Query {
	if q.err != nil {
		return q
	}
	q.offset = n
	return q
}

// FindOne 查询单条记录
func (q *Query) FindOne(itf interface{}) *Query {
	q.Prepare()
	if q.err != nil {
		return q
	}

	rows, err := q.db.Query(q.query, q.args...)
	if err != nil {
		q.err = err
		return q
	}

	if rows != nil {
		defer func() {
			rows.Close()
		}()
	}

	if !rows.Next() {
		q.err = ErrEmptyResult
		return q
	}

	columns, err := rows.Columns()
	if err != nil {
		q.err = err
		return q
	}

	dest, err := unmarshalStruct(itf, columns)
	if err != nil {
		q.err = err
		return q
	}

	q.err = rows.Scan(dest...)

	return q
}

// Find 查询多条记录
func (q *Query) Find(itf interface{}) *Query {
	q.Prepare()
	if q.err != nil {
		return q
	}

	sliceValue := reflect.Indirect(reflect.ValueOf(itf))
	if sliceValue.Kind() != reflect.Slice {
		q.err = ErrNotSlice
		return q
	}

	t := sliceValue.Type().Elem()
	k := t.Kind()
	if k == reflect.Ptr {
		t = t.Elem()
	}

	rows, err := q.db.Query(q.query, q.args...)
	if err != nil {
		q.err = err
		return q
	}

	if rows != nil {
		defer func() {
			rows.Close()
		}()
	}

	columns, err := rows.Columns()
	if err != nil {
		q.err = err
		return q
	}

	for rows.Next() {
		v := reflect.New(t)
		dest, err := unmarshalValue(reflect.Indirect(v), columns)
		if err != nil {
			q.err = err
			return q
		}

		err = rows.Scan(dest...)
		if err != nil {
			q.err = err
			return q
		}
		if k != reflect.Ptr {
			v = v.Elem()
		}

		sliceValue.Set(reflect.Append(sliceValue, v))
	}

	return q
}
