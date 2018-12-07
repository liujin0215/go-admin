package response

import (
	"encoding/json"

	"github.com/kataras/iris"
)

//Response 通用回复结构体
type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data,omitempty"`
	Err  string      `json:"err,omitempty"`
}

//WriteResp 通用回复方法
func WriteResp(ctx iris.Context, resp *Response) {
	json.NewEncoder(ctx.ResponseWriter()).Encode(resp)
}

func WriteSuccessResp(ctx iris.Context, data interface{}) {
	WriteResp(ctx, &Response{Data: data})
}

func WriteFailResp(ctx iris.Context, errNo int) {
	WriteResp(ctx, &Response{Err: errMap[errNo]})
}
