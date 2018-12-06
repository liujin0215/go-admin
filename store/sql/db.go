package sql

import (
	"bytes"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
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
	sQuery, args := prepareMapQuery(model, ",")
	if len(args) == 0 {
		return 0, ErrQueryEmpty
	}
	wQuery, wArgs := prepareWhere(where, wArgs)
	if len(wArgs) == 0 {
		return 0, ErrConditionEmpty
	}
	args = append(args, wArgs)
	var result sql.Result
	query := fmt.Sprintf("update %s set %s where %s;", tbName, sQuery, wQuery)
	result, err = db.DB.Exec(query, args...)
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

func prepareSelect(itf interface{}) string {
	t := reflect.TypeOf(itf)
	v := reflect.ValueOf(itf)

	if t.Kind() == reflect.String {
		return itf.(string)
	}

	if t.Kind() == reflect.Ptr {
		v = v.Elem()
		t = v.Type()
	}

	if t.Kind() != reflect.Struct {
		return ""
	}

	var buff bytes.Buffer
	for i := 0; i < t.NumField(); i++ {
		key := getFieldName(t.Field(i))
		if len(key) == 0 {
			continue
		}
		buff.WriteString(key)
		if i < t.NumField()-1 {
			buff.WriteString(",")
		}
	}
	return buff.String()
}

func getFieldName(field reflect.StructField) string {
	tag := field.Tag.Get(tagName)
	if len(tag) == 0 {
		return ""
	}
	options := strings.Split(tag, ",")
	return strings.TrimSpace(options[0])
}

func prepareWhere(where interface{}, args []interface{}) (string, []interface{}) {
	if query, ok := where.(string); ok {
		return query, args
	}

	if model, ok := where.(MapModel); ok {
		return prepareMapQuery(model, " and ")
	}

	return "", nil
}

func parseMap(model map[string]interface{}) (kList []string, vList []interface{}) {
	for key, value := range model {
		kList = append(kList, key)
		vList = append(vList, value)
	}
	return
}

func prepareMapQuery(model map[string]interface{}, sep string) (query string, args []interface{}) {
	var kList []string
	kList, args = parseMap(model)
	query = joinKlist(kList, sep)
	return
}

func joinKlist(kList []string, sep string) string {
	l := len(kList)
	var buff bytes.Buffer
	for n, key := range kList {
		buff.WriteString(key)
		buff.WriteString("=?")
		if n < l-1 {
			buff.WriteString(sep)
		}
	}
	return buff.String()
}
