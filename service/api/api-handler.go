package api


import (
	"net/http"
)

func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.GET("/", rt.getHelloWorld)
	rt.router.GET("/context", rt.wrap(rt.getContextReply))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)
	rt.router.POST("/session", rt.wrap(doLogin))
	return rt.router

}


