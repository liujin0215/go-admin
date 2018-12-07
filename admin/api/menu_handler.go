package api

import (
	"encoding/json"
	"fmt"
	"go-admin/admin/model"
	"go-admin/admin/utils"
	"go-admin/response"
	"go-admin/store/sql"

	"github.com/kataras/iris"
)

//CreateMenu 创建菜单
func CreateMenu(ctx iris.Context) {
	form := make(sql.MapModel)
	err := json.NewDecoder(ctx.Request().Body).Decode(&form)
	if err != nil {
		response.WriteFailResp(ctx, response.ErrParseBody)
		return
	}

	id, err := utils.AdminDB.Insert(model.TbMenu, form)
	if err != nil {
		response.WriteFailResp(ctx, response.ErrCreate)
		return
	}

	response.WriteSuccessResp(ctx, id)
	return
}

//RetrieveMenu 查询菜单
func RetrieveMenu(ctx iris.Context) {
	form := make(sql.MapModel)
	err := json.NewDecoder(ctx.Request().Body).Decode(&form)
	if err != nil {
		response.WriteFailResp(ctx, response.ErrParseBody)
		return
	}

	fmt.Println(form)

	var records []*model.Menu
	err = utils.AdminDB.Tb(model.TbMenu).Select(new(model.Menu)).Where(form).Find(&records).Err()
	if err != nil {
		fmt.Println(err)
		response.WriteFailResp(ctx, response.ErrRetrieve)
		return
	}

	response.WriteSuccessResp(ctx, &records)
	return
}

//UpdateMenu 更新菜单
func UpdateMenu(ctx iris.Context) {
	form := make(sql.MapModel)
	err := json.NewDecoder(ctx.Request().Body).Decode(&form)
	if err != nil {
		response.WriteFailResp(ctx, response.ErrParseBody)
		return
	}

	if form.PK() == 0 {
		response.WriteFailResp(ctx, response.ErrEmptyPK)
		return
	}

	affected, err := utils.AdminDB.Update(model.TbMenu, form, sql.MapModel{"id": form.PK()})
	if err != nil {
		fmt.Println(err)
		response.WriteFailResp(ctx, response.ErrCreate)
		return
	}

	response.WriteSuccessResp(ctx, affected)
	return
}

//DeleteMenu 删除菜单
func DeleteMenu(ctx iris.Context) {
	form := make(sql.MapModel)
	err := json.NewDecoder(ctx.Request().Body).Decode(&form)
	if err != nil {
		response.WriteFailResp(ctx, response.ErrParseBody)
		return
	}

	if form.PK() == 0 {
		response.WriteFailResp(ctx, response.ErrEmptyPK)
		return
	}

	affected, err := utils.AdminDB.Delete(model.TbMenu, sql.MapModel{"id": form.PK()})
	if err != nil {
		response.WriteFailResp(ctx, response.ErrCreate)
		return
	}

	response.WriteSuccessResp(ctx, affected)
	return
}
