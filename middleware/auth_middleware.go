package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/valyala/fasthttp"
	"gitlab.com/visaruttirataworawan/grader_gw/interfaces"
	"gitlab.com/visaruttirataworawan/grader_gw/utils"
	"strings"
	"time"
)

func JWT(tokenInput string) (bool, *string) {
	type MyCustomClaims struct {
		Sub  string    `json:"sub"`
		Name string    `json:"name"`
		exp  time.Time `json:"exp"`
		iat  time.Time `json:"iat"`
		jwt.MapClaims

	}
	verificationKey := "my_secret_key"
	tkn, err := jwt.ParseWithClaims(tokenInput, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(verificationKey), nil
	})
	if err != nil {
		if tkn == nil {
			message := "Bad request"
			return false, &message
		}
		if !tkn.Valid {
			message := "Unauthorized"
			return false, &message
		}
	}

	if claims, ok := tkn.Claims.(*MyCustomClaims); ok && tkn.Valid {
		fmt.Printf("Hello %v will expired at %v", claims.Name, claims.exp)
		return true, nil
	}
	message := "Bad request"
	return false, &message
}

func AuthMiddleware(h fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		header := string(ctx.Request.Header.Peek("Authorization"))
		var token string
		if strings.HasPrefix(header, "Bearer ") {
			token = header[7:]
		}
		result, err := JWT(token)
		if result == false && *err == "Bad request" {
			res := interfaces.CreateHTTPResponsePayload(nil, "Failed", "Bad Request")
			utils.HTTPResponse(ctx, res, 400)
			//ctx.Error(*err, fasthttp.StatusBadRequest)
		} else if result == false && *err == "Unauthorized" {
			res := interfaces.CreateHTTPResponsePayload(nil, "Failed", "Unauthorized")
			utils.HTTPResponse(ctx, res, 401)
		} else {
			h(ctx)
			return
		}
	}
}
