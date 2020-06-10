package utils

import (
	"encoding/json"
	// "fmt"
	// "net/http"

	// routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

// var status interfaces.AckStatus


func HTTPResponse(ctx *fasthttp.RequestCtx, response interface{}, statusCode int) error {
	responseBytes, err := json.Marshal(response)
	if err != nil {
		return err
	}

	_, err = ctx.Write(responseBytes)
	ctx.SetStatusCode(statusCode)
	ctx.SetContentType("application/json")
	if err != nil {
		return err
	}

	ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
	//ctx.Response.Header.Set("Access-Control-Allow-Origin", "http://128.199.87.22:3000")
	//ctx.Response.Header.Set("Access-Control-Allow-Origin", "http://localhost:3000")
	ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
	ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET,HEAD,PUT,POST,DELETE")
	ctx.Response.Header.Set("Access-Control-Allow-Headers", "Origin,Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Set-Cookie")

	// if string(ctx.Method()) == "OPTIONS" {
	// 	ctx.Abort()
	// }

	// if err := ctx.Next(); err != nil {
	// 	if httpError, ok := err.(routing.HTTPError); ok {
	// 		ctx.Response.SetStatusCode(httpError.StatusCode())
	// 	} else {
	// 		ctx.Response.SetStatusCode(http.StatusInternalServerError)
	// 	}



		
	// }
	// b, _ := json.Marshal(&JSONError{
	// 	Message: fmt.Sprintf("%s", err),
	// })
	// ctx.SetContentType("application/json")
	// ctx.SetBody(b)
	return nil

}
