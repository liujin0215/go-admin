package crud

import (
	"encoding/json"
	"go-admin/response"
	"go-admin/store/sql"

	"github.com/kataras/iris"
)

// NewCRUD 生成crud的party handler
func NewCRUD(app iris.Party, route string, db *sql.DB, model sql.Model) iris.Party {
	return app.PartyFunc(route, func(u iris.Party) {
		u.Post("/create", NewHandler(db, sql.InsertMethod, model))
		u.Post("/update", NewHandler(db, sql.UpdateMethod, model))
		u.Post("/delete", NewHandler(db, sql.DeleteMethod, model))
		u.Get("/retrieve", NewHandler(db, sql.SelectMethod, model))
		u.Get("/retrieveone", NewHandler(db, sql.SelectOneMethod, model))
	})
}

// NewHandler 生成新的handler
func NewHandler(db *sql.DB, method int, model sql.Model) iris.Handler {
	switch method {
	case sql.InsertMethod, sql.UpdateMethod, sql.DeleteMethod:
		return func(ctx iris.Context) {
			form := make(sql.MapModel)
			err := json.NewDecoder(ctx.Request().Body).Decode(&form)
			if err != nil {
				response.WriteFailResp(ctx, response.ErrParseBody)
				return
			}
			ms := &sql.ModelStruct{
				MapModel: form,
				Model:    model,
			}

			_, err = db.Exec(method, ms)
			if err != nil {
				response.WriteFailResp(ctx, method)
				return
			}

			response.WriteSuccessResp(ctx, nil)
			ctx.Next()
			return
		}
	case sql.SelectMethod:
		return func(ctx iris.Context) {
			var err error
			form := make(sql.MapModel)
			err = json.NewDecoder(ctx.Request().Body).Decode(&form)
			if err != nil {
				response.WriteFailResp(ctx, response.ErrParseBody)
				return
			}
			ms := &sql.ModelStruct{
				MapModel: form,
				Model:    model,
			}

			resp := new(RetrieveData)

			err = db.Tb(model.TbName()).Select("count(*)").Where(form).FindOne(&resp.Count).Err()
			if err != nil || resp.Count == 0 {
				response.WriteFailResp(ctx, response.ErrRetrieve)
				return
			}

			res, err := db.Retrieve(ms)
			if err != nil {
				response.WriteFailResp(ctx, response.ErrRetrieve)
				return
			}
			resp.Record = res

			response.WriteSuccessResp(ctx, resp)
			ctx.Next()
			return
		}
	case sql.SelectOneMethod:
		return func(ctx iris.Context) {
			var err error
			form := make(sql.MapModel)
			err = json.NewDecoder(ctx.Request().Body).Decode(&form)
			if err != nil {
				response.WriteFailResp(ctx, response.ErrParseBody)
				return
			}
			ms := &sql.ModelStruct{
				MapModel: form,
				Model:    model,
			}

			res, err := db.RetrieveOneEx(ms)
			if err != nil {
				response.WriteFailResp(ctx, response.ErrRetrieve)
				return
			}

			response.WriteSuccessResp(ctx, res)
			ctx.Next()
			return
		}
	default:
		return nil
	}
}
