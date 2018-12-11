package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-admin/admin/utils"
	"go-admin/store/sql"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func initDB() (err error) {
	cfg := &sql.DBConfig{
		Addr: "root:Liujin@tcp(localhost:3306)/data?charset=utf8mb4&parseTime=true",
		Type: "mysql",
	}

	utils.AdminDB, err = sql.NewDB(cfg)
	return
}

func test(t *testing.T, handler context.Handler, data interface{}) {
	err := initDB()
	if err != nil {
		t.Fatal(err)
	}
	byteData, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}
	r := httptest.NewRequest("POST", "/", bytes.NewReader(byteData))
	w := httptest.NewRecorder()

	app := iris.New()
	ctx := context.NewContext(app)
	ctx.BeginRequest(w, r)
	handler(ctx)
	ctx.EndRequest()

	resp := w.Result()
	defer resp.Body.Close()
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(respData))
}

func TestCreateMenu(t *testing.T) {
	data := sql.MapModel{"name": "test"}
	test(t, CreateMenu, data)
}

func TestRetrieveMenu(t *testing.T) {
	data := sql.MapModel{"id": 1}
	test(t, RetrieveMenu, data)
}

func TestUpdateMenu(t *testing.T) {
	data := sql.MapModel{"id": 1, "name": "菜单配置"}
	test(t, UpdateMenu, data)
}

func TestDeleteMenu(t *testing.T) {
	data := sql.MapModel{"id": 1}
	test(t, DeleteMenu, data)
}
