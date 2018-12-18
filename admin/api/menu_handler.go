package api

import (
	"encoding/json"
	"fmt"
	"go-admin/admin/model"
	"go-admin/admin/utils"
	"go-admin/headers/controller"
	"go-admin/response"
	"go-admin/store/sql"

	"github.com/kataras/iris"
)

func RegisterRouteForMenu(party iris.Party) {
	menu := party.Party("/menu", controller.AccessHeaders)
	menu.AllowMethods("Post", "Options")
	menu.Handle("POST", "/create", CreateMenu)
	menu.Handle("POST", "/retrieve", RetrieveMenu)
	menu.Handle("POST", "/update", UpdateMenu)
	menu.Handle("POST", "/delete", DeleteMenu)
}

func RegisterPermissionForMenu(party iris.Party) {
	//basePath := menu.GetRelPath

}

//CreateMenu 创建菜单
func CreateMenu(ctx iris.Context) {
	form := make(sql.MapModel)
	err := json.NewDecoder(ctx.Request().Body).Decode(&form)
	if err != nil {
		response.WriteFailResp(ctx, response.ErrParseBody)
		return
	}
	baseModel := new(model.Menu)

	res, err := utils.AdminDB.Exec(sql.InsertMethod, &sql.ModelStruct{form, baseModel})
	if err != nil {
		response.WriteFailResp(ctx, response.ErrCreate)
		return
	}
	id, err := res.LastInsertId()
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

	baseModel := new(model.Menu)
	var records []*model.Menu
	err = utils.AdminDB.Tb(model.TbMenu).Select(baseModel).Where(form, baseModel).Find(&records).Err()
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

	baseModel := new(model.Menu)

	res, err := utils.AdminDB.Exec(sql.UpdateMethod, &sql.ModelStruct{form, baseModel})
	if err != nil {
		response.WriteFailResp(ctx, response.ErrCreate)
		return
	}
	affected, err := res.RowsAffected()
	if err != nil {
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

	baseModel := new(model.Menu)

	res, err := utils.AdminDB.Exec(sql.DeleteMethod, &sql.ModelStruct{form, baseModel})
	if err != nil {
		response.WriteFailResp(ctx, response.ErrCreate)
		return
	}
	affected, err := res.RowsAffected()
	if err != nil {
		response.WriteFailResp(ctx, response.ErrCreate)
		return
	}

	response.WriteSuccessResp(ctx, affected)
	return
}
