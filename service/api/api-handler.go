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
	rt.router.POST("/session", rt.wrap(rt.doLogin)) // to vheck if it ckecks if user is logged
	//user
	rt.router.PUT("/session/:userId/userName", rt.wrap(rt.setMyUsername))
	
	rt.router.PUT("/session/:userId/userPhoto", rt.wrap(rt.setMyPhoto))

	//MESSAGES
	rt.router.PUT("/message", rt.wrap(rt.sendMessage))
	rt.router.DELETE("/message/:messageId", rt.wrap(rt.deleteMessage))
	rt.router.PUT("/messsage/forward/:messageId/:chatId", rt.wrap(rt.forwardMessage))
	
	return rt.router

}


