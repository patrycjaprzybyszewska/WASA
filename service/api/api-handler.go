package api

import (
	"net/http"
)

func (rt *_router) Handler() http.Handler {

	rt.router.GET("/", rt.getHelloWorld)
	rt.router.GET("/context", rt.wrap(rt.getContextReply))

	rt.router.GET("/liveness", rt.liveness)
	rt.router.POST("/session", rt.wrap(rt.doLogin)) // to vheck if it ckecks if user is logged
	// user
	rt.router.PUT("/session/:userId/userName", rt.wrap(rt.setMyUserName))
	rt.router.PUT("/session/:userId/userPhoto", rt.wrap(rt.setMyPhoto))

	// MESSAGES
	rt.router.PUT("/message", rt.wrap(rt.sendMessage))
	rt.router.DELETE("/message/:messageId", rt.wrap(rt.deleteMessage))
	rt.router.PUT("/message/forward/:messageId/:chatId", rt.wrap(rt.forwardMessage))
	rt.router.PUT("/message/comment/:messageId", rt.wrap(rt.commentMessage))
	rt.router.DELETE("/comment/:commentId", rt.wrap(rt.uncommentMessage))

	rt.router.GET("/conversation/:chatId", rt.wrap(rt.getConversation))

	// chat
	rt.router.PUT("/groupchat/:chatId/add/:userId", rt.wrap(rt.addToGroup))
	rt.router.DELETE("/groupchat/:chatId/leave/:userId", rt.wrap(rt.leaveGroup))
	rt.router.PUT("/groupchat/:chatId/groupName", rt.wrap(rt.setGroupName))
	rt.router.PUT("/groupchat/:chatId/groupPhoto", rt.wrap(rt.setGroupPhoto))
	rt.router.GET("/conversation", rt.wrap(rt.getConversations))

	return rt.router

}
