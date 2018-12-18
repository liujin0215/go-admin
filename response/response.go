package response

import (
	"github.com/kataras/iris"
)

//Response 通用回复结构体
type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data,omitempty"`
	Err  string      `json:"err,omitempty"`
}

//WriteResp 通用回复
func WriteResp(ctx iris.Context, resp *Response) {
	ctx.JSON(resp)
	//ctx.NotFound()
}

//WriteSuccessResp 成功回复
func WriteSuccessResp(ctx iris.Context, data interface{}) {
	WriteResp(ctx, &Response{Data: data})
}

//WriteFailResp 失败回复
func WriteFailResp(ctx iris.Context, errNo int) {
	WriteResp(ctx, &Response{Err: errMap[errNo]})
}
