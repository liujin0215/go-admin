package admin

import (
	"go-admin/admin/model"
	"go-admin/admin/utils"
	"go-admin/crud"

	"github.com/kataras/iris"
)

// RegisterRoute 注册路由
func RegisterRoute(app iris.Party) {
	crud.NewCRUD(app, "/menu", utils.AdminDB, new(model.Menu))
	crud.NewCRUD(app, "/role", utils.AdminDB, new(model.Role))
	crud.NewCRUD(app, "/adminuser", utils.AdminDB, new(model.AdminUser))
	crud.NewCRUD(app, "/permission", utils.AdminDB, new(model.Permission))
}
