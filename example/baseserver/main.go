package main

import (
	"flag"
	"go-admin/admin"
	"go-admin/example/baseserver/config"
	hcontroller "go-admin/headers/controller"

	"github.com/kataras/iris"
)

var path = flag.String("file", "./etc/config.json", "config file")

func main() {
	err := config.LoadConfig(*path)
	if err != nil {
		panic(err)
	}

	app := iris.Default()
	app.Use(hcontroller.AccessHeaders)

	admin.RegisterRoute(app)

	app.Run(iris.Addr(config.Conf.Addr))
}
