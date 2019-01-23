//Package sql ModelStruct类
package sql

import "fmt"

// ModelStruct ModelStruct类
type ModelStruct struct {
	MapModel
	Model
}

// Prepare 解析字段，并整合sql语句
func (ms *ModelStruct) Prepare(method int) StmtStruct {
	switch method {
	case InsertMethod:
		return ms.prepareInsert()
	case UpdateMethod:
		return ms.prepareUpdate()
	case DeleteMethod:
		return ms.prepareDelete()
	case SelectMethod, SelectOneMethod:
		return ms.prepareSelect()
	default:
		return nil
	}
}

// prepare 解析字段
func (ms *ModelStruct) prepare() (keyList []string, valueList []interface{}) {
	if ms.MapModel == nil {
		keyList, valueList = parseStruct(ms.Model)
	} else {
		keyList, valueList = ms.MapModel.Parse(ms.Model)
	}
	return
}

// PK 解析主键
func (ms *ModelStruct) PK() (id uint) {
	if ms.MapModel == nil {
		id = ms.Model.PK()
	} else {
		id = ms.MapModel.PK()
	}
	return
}

// prepareInsert 解析并准备insert语句
func (ms *ModelStruct) prepareInsert() StmtStruct {
	keyList, valueList := ms.prepare()
	return &StringStruct{
		stmt: fmt.Sprintf("insert into %s set %s;", ms.Model.TbName(), joinKlist(keyList, ",")),
		args: valueList,
	}
}

// prepareUpdate 解析并准备update语句
func (ms *ModelStruct) prepareUpdate() StmtStruct {
	keyList, valueList := ms.prepare()
	valueList = append(valueList, ms.PK())
	return &StringStruct{
		stmt: fmt.Sprintf("update %s set %s where id = ?;", ms.Model.TbName(), joinKlist(keyList, ",")),
		args: valueList,
	}
}

// prepareDelete 解析并准备delete语句
func (ms *ModelStruct) prepareDelete() StmtStruct {
	keyList, valueList := ms.prepare()
	return &StringStruct{
		stmt: fmt.Sprintf("delete from %s where %s;", ms.Model.TbName(), joinKlist(keyList, ",")),
		args: valueList,
	}
}

// prepareSelect 解析并准备select语句
func (ms *ModelStruct) prepareSelect() StmtStruct {
	keyList, valueList := ms.MapModel.Parse(ms.Model)
	ss := &StringStruct{
		stmt: fmt.Sprintf("select %s from %s", prepareSelect(ms.Model), ms.TbName()),
		args: valueList,
	}

	if len(valueList) > 0 {
		ss.stmt += fmt.Sprintf(" where %s", joinKlist(keyList, " and "))
	}

	if len(ms.OrderBy()) > 0 {
		ss.stmt += fmt.Sprintf(" order by %s", ms.OrderBy())
	}

	limit := ms.Limit()
	if limit > 0 {
		ss.stmt += " limit ?"
		ss.args = append(ss.args, limit)
	}

	offset := ms.Offset()
	if offset > 0 {
		ss.stmt += " offset ?"
		ss.args = append(ss.args, offset)
	}

	ss.stmt += ";"
	return ss
}
