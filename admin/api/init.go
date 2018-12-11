package api

import (
	"go-admin/headers/controller"

	"github.com/kataras/iris"
)

//RegisterRoute 注册路由
func RegisterRoute(app *iris.Application) {
	app.HandleMany("POST OPTIONS", "/api/menu/create", controller.AccessHeaders, CreateMenu)
	app.HandleMany("POST OPTIONS", "/api/menu/retrieve", controller.AccessHeaders, RetrieveMenu)
	app.HandleMany("POST OPTIONS", "/api/menu/update", controller.AccessHeaders, UpdateMenu)
	app.HandleMany("POST OPTIONS", "/api/menu/delete", controller.AccessHeaders, DeleteMenu)
}
