package api


import (
	"net/http"
)

func (rt *_router) Handler() http. Handler {
	rt.router.POST("/message", rt.sendMessage)
	rt.router.DELETE("/message/:id", rt.deleteMessage)
	rt.routrt.PUT("/message/:id", rt.commentMessage)
}