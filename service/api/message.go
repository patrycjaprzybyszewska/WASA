package api

import (
	// "git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"encoding/json"
	"fmt"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"time"
)

func (rt *_router) sendMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")

	var message Message
	var requestBody struct {
		Content  string `json:"content"`
		ChatName string `json:"chatName"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	message.SenderId, err = auth(r.Header.Get("Authorization"))
	if err != nil {
		http.Error(w, "Unauthorized access", http.StatusUnauthorized)
		return
	}

	message.SenderName, err = rt.db.GetUserNameById(message.SenderId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if requestBody.Content == "" || requestBody.ChatName == "" {
		http.Error(w, "Message cannot be sent, missing informations", http.StatusBadRequest)
		return
	}

	message.ChatId, err = rt.db.GetChatIdbyName(requestBody.ChatName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	message.Content = requestBody.Content
	currentTime := time.Now()
	message.MessageDate = currentTime.Format("2006-01-02")
	message.MessageTime = currentTime.Format("15:04")
	message.State = "delivered"

	dbmessage, err := rt.db.Sendmessage(message.MessageToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	message.MessageFromDatabase(dbmessage)
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(message)

	// 201 send correctly, 400 missing info potzrebuje user id, DODAnC AUTORYZACJE
}

func (rt *_router) deleteMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Missing authorization", http.StatusUnauthorized)
		return
	}
	messageId, err := strconv.ParseInt(ps.ByName("messageId"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid message ID", http.StatusBadRequest)
		return
	}

	err = rt.db.Removemessage(uint64(messageId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
} // 204 no content 404, i need mess id, add auth

func (rt *_router) forwardMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var message Message

	messageId, err := strconv.ParseUint(ps.ByName("messageId"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid message ID", http.StatusBadRequest)
		return
	}

	dbmessage, err := rt.db.GetMessageById(messageId)
	if err != nil {
		http.Error(w, "Forwarded message does not exist", http.StatusNotFound)
		return
	}
	var requestBody struct {
		ChatName string `json:"chatName"`
	}
	err = json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if requestBody.ChatName == "" {
		http.Error(w, "Message cannot be sent, missing informations", http.StatusBadRequest)
		return
	}
	message.ChatId, err = rt.db.GetChatIdbyName(requestBody.ChatName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Missing authorization", http.StatusUnauthorized)
		return
	}
	message.SenderId, err = auth(authHeader)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}
	message.SenderName, err = rt.db.GetUserNameById(message.SenderId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	currentTime := time.Now()
	message.MessageDate = currentTime.Format("2006-01-02")
	message.MessageTime = currentTime.Format("15:04")
	message.State = "send"
	if dbmessage.Content == "" {
		http.Error(w, "Message content is empty", http.StatusBadRequest)
		return
	}
	message.Content = dbmessage.Content
	dbmessage, err = rt.db.Sendmessage(message.MessageToDatabase())
	if err != nil {
		http.Error(w, "Error forwarding message", http.StatusInternalServerError)
		return
	}

	message.MessageFromDatabase(dbmessage)
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(message)
}

// 201 404

func (rt *_router) commentMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Missing authorization", http.StatusUnauthorized)
		return
	}
	var comment Comment
	messageId, err := strconv.ParseUint(ps.ByName("messageId"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid message ID", http.StatusBadRequest)
		return
	}
	err = rt.db.CheckMessageById(messageId)
	if err != nil {
		http.Error(w, "Message to comment not found", http.StatusNotFound)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		http.Error(w, "Invalid or empty comment content", http.StatusBadRequest)
		return
	}

	dbcomment, err := rt.db.Commentmessage(comment.CommentToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	comment.CommentFromDatabase(dbcomment)
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(comment)

} // 201 404

func (rt *_router) uncommentMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Missing authorization", http.StatusUnauthorized)
		return
	}
	commentId, err := strconv.ParseInt(ps.ByName("commentId"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid comment ID", http.StatusBadRequest)
		return
	}
	// sprawdzic czy kom istnieje
	err = rt.db.CheckCommentById(uint64(commentId))
	if err != nil {
		http.Error(w, "Comment not found", http.StatusNotFound)
		return
	}

	err = rt.db.Removecomment(uint64(commentId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
} // 204 404

func (rt *_router) getConversation(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Missing authorization", http.StatusUnauthorized)
		return
	}
	chatIdStr := ps.ByName("chatId")
	chatId, err := strconv.ParseUint(chatIdStr, 10, 64)

	conversation, err := rt.db.GetConversation(chatId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(conversation)
}
