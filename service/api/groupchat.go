package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (rt *_router) addToGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Missing authorization", http.StatusUnauthorized)
		return
	}
	chatIdStr := ps.ByName("chatId")
	chatId, err := strconv.ParseUint(chatIdStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid chat ID", http.StatusBadRequest)
		return
	}

	userName := ps.ByName("userId")
	userId, err := rt.db.GetUserIdByName(userName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = rt.db.AddUserToChat(chatId, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Chat or user not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte(`{"message": "User ID has been added sucesfullly"}`))
}

func (rt *_router) leaveGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Missing authorization", http.StatusUnauthorized)
		return
	}
	chatId, err := strconv.ParseUint(ps.ByName("chatId"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid chat ID", http.StatusBadRequest)
		return
	}

	userIdStr := ps.ByName("userId")
	userId, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	err = rt.db.LeaveGroup(chatId, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Chat or user not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte(`{"message": "Chat left properly"}`))
}

func (rt *_router) setGroupName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Missing authorization", http.StatusUnauthorized)
		return
	}

	var chat Chat
	err := json.NewDecoder(r.Body).Decode(&chat)
	if err != nil {
		http.Error(w, "Invalid request body ", http.StatusBadRequest)
		return
	}
	chat.ChatId, err = strconv.ParseUint(ps.ByName("chatId"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid chat ID", http.StatusBadRequest)
		return
	}

	chat.ChatPhoto, err = rt.db.GetChatPhotoById(chat.ChatId)
	if err != nil {
		http.Error(w, "No photo", http.StatusInternalServerError)
		return
	}

	dbchat, err := rt.db.SetGroupName(chat.ChatToDatabase(), chat.ChatName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Chat or user not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	chat.ChatFromDatabase(dbchat)
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(chat)
}
func (rt *_router) setGroupPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Missing authorization", http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var chat Chat
	err := json.NewDecoder(r.Body).Decode(&chat)
	if err != nil {
		http.Error(w, "Invalid request body ", http.StatusBadRequest)
		return
	}
	chat.ChatId, err = strconv.ParseUint(ps.ByName("chatId"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid chat ID", http.StatusBadRequest)
		return
	}

	chat.ChatName, err = rt.db.GetChatNameById(chat.ChatId)
	if err != nil {
		http.Error(w, "No chat name", http.StatusInternalServerError)
		return
	}

	dbchat, err := rt.db.SetGroupName(chat.ChatToDatabase(), chat.ChatPhoto)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Chat or user not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	chat.ChatFromDatabase(dbchat)
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(chat)
}

func (rt *_router) getConversations(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")
	userId, err := auth(r.Header.Get("Authorization"))
	if err != nil {
		http.Error(w, " access", http.StatusUnauthorized)
		return
	}

	chats, err := rt.db.GetChats(userId)
	if err != nil {
		http.Error(w, "Unable to fetch conversations", http.StatusInternalServerError)
		return
	}

	if len(chats) == 0 {
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode([]Chat{})
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(chats)
}
