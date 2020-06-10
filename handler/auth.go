package handler

import (
	"crypto/tls"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"gitlab.com/visaruttirataworawan/grader_gw/utils"
	"log"
)

func Signin(ctx *fasthttp.RequestCtx) {

	authServer := utils.GetConfig("AUTH_HOST")
	req := &ctx.Request
	req.SetRequestURI(authServer + "/signin")

	res := &fasthttp.Response{}
	clientL := &fasthttp.Client{
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	if err := clientL.Do(req, res); err != nil {
		log.Println("get response failed", err)
	}

	resBodyByte := res.Body()
	var bodyRes map[string]interface{}
	json.Unmarshal(resBodyByte, &bodyRes)
	utils.HTTPResponse(ctx, bodyRes, res.Header.StatusCode())
}
