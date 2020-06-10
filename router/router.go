package router

import ( routing "github.com/buaazp/fasthttprouter"
)
func New() *routing.Router{
	return routing.New()
}

func Mount(r routing.Router){
	//api gateway forward to auth

}