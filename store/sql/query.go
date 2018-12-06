package sql

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
)

type (
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

func (q *Query) Err() error {
	return q.err
}

func (q *Query) prepare() {
	if q.err != nil {
		return
	}

	if len(q.query) > 0 {
		return
	}

	if len(q.tb) == 0 || len(q.selectQuery) == 0 {
		q.err = ErrQueryEmpty
		return
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
	fmt.Println(q.query)
}

func (q *Query) Tb(tbName string) *Query {
	if q.err != nil {
		return q
	}
	q.tb = tbName
	return q
}

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

func (q *Query) Where(where interface{}, args ...interface{}) *Query {
	if q.err != nil {
		return q
	}
	q.where, q.args = prepareWhere(where, args)
	return q
}

func (q *Query) OrderBy(args ...string) *Query {
	if q.err != nil {
		return q
	}
	q.orderby = append(q.orderby, args...)
	return q
}

func (q *Query) Limit(n uint) *Query {
	if q.err != nil {
		return q
	}
	q.limit = n
	return q
}

func (q *Query) Offset(n uint) *Query {
	if q.err != nil {
		return q
	}
	q.offset = n
	return q
}

func (q *Query) FindOne(itf interface{}) *Query {
	q.prepare()
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

func (q *Query) Find(itf interface{}) *Query {
	q.prepare()
	if q.err != nil {
		return q
	}

	sliceValue := reflect.Indirect(reflect.ValueOf(itf))
	if sliceValue.Kind() != reflect.Slice {
		q.err = ErrNotSlice
		return q
	}

	t := sliceValue.Type().Elem()
	if t.Kind() != reflect.Ptr {
		q.err = ErrNotPtr
		return q
	}

	t = t.Elem()

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

		sliceValue.Set(reflect.Append(sliceValue, v.Elem().Addr()))
	}

	return q
}
