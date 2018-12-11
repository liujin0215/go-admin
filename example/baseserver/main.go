package main

import (
	"flag"
	"go-admin/admin/api"
	"go-admin/example/baseserver/config"

	"github.com/kataras/iris"
)

var path = flag.String("file", "./etc/config.json", "config file")

func main() {
	err := config.LoadConfig(*path)
	if err != nil {
		panic(err)
	}

	app := iris.Default()
	api.RegisterRoute(app)

	app.Run(iris.Addr(config.Conf.Addr))
}
