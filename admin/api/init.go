package api

import (
	"github.com/kataras/iris"
)

//RegisterRoute 注册路由
func RegisterRoute(app *iris.Application) {
	app.HandleMany("POST OPTIONS", "/menu/create", CreateMenu)
	app.HandleMany("POST OPTIONS", "/menu/retrieve", RetrieveMenu)
	app.HandleMany("POST OPTIONS", "/menu/update", UpdateMenu)
	app.HandleMany("POST OPTIONS", "/menu/delete", DeleteMenu)
}
