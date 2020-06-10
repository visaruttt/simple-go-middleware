package handler

import (
	"github.com/valyala/fasthttp"
	"gitlab.com/visaruttirataworawan/grader_gw/interfaces"
	"gitlab.com/visaruttirataworawan/grader_gw/utils"
)

func Helloworld(ctx *fasthttp.RequestCtx){
	res := interfaces.CreateHTTPResponsePayload("helloworld","Success",nil)
	utils.HTTPResponse(ctx, res,200)
}
