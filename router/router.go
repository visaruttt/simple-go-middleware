package router

import (
	routing "github.com/buaazp/fasthttprouter"
	"gitlab.com/visaruttirataworawan/grader_gw/handler"
	"gitlab.com/visaruttirataworawan/grader_gw/middleware"
)

func New() *routing.Router {
	return routing.New()
}

func Mount(r *routing.Router) {
	//api gateway forward to auth
	r.POST("/signin", handler.Signin)
	//router.GET("/protected/", middleware.AuthMiddleware(handler.Protected))
	r.GET("/helloworld", middleware.AuthMiddleware(handler.Helloworld))

}
