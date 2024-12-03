package api


import (
	"net/http"
)

func (rt *_router) Handler() http. Handler {
	rt.router.POST("/message", rt.sendMessage)
	rt.router.DELETE("/message/:id", rt.deleteMessage)
	rt.routrt.PUT("/message/:id", rt.commentMessage)
	rt.router.POST("/message/:id", rt.forwardMessage)
	rt.router.DELETE("/message/comment/:id", rt.deleteComment)
	rt.router.POST("/session", rt.doLogin)
	rt.router.PUT("session/:id/userName", rt.setMyUsername)
}


