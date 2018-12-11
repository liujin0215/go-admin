package controller

import "github.com/kataras/iris"

//AccessHeaders 跨域请求头设置
func AccessHeaders(ctx iris.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "OPTIONS POST")
	ctx.Header("Access-Control-Allow-Headers", "content-type")

	if ctx.Method() == "OPTIONS" {
		return
	}
	ctx.Next()
}
