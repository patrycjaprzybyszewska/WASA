package api


import (
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"net/http"
	"encoding/json"
	"strconv"
	"time"
	"fmt"
)

func (rt *_router) sendMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")
	var message Message

	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	
	if message.Content == "" ||  message.ChatId == 0 {
		http.Error(w, "Missing required fields: content or chatId", http.StatusBadRequest)
		return
	}

	currentTime := time.Now()
	message.MessageDate = currentTime.Format("2006-01-02") // Format YYYY-MM-DD
	message.MessageTime = currentTime.Format("15:04")      // Format HH:MM
	message.State = "delivered"

	dbmessage, err := rt.db.Sendmessage(message.MessageToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	

	message.MessageFromDatabase(dbmessage)
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(message)
////201 send correctly, 400 missing info potzrebuje user id, DODAC AUTORYZACJE
}

func (rt *_router) deleteMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	
	messageIdStr := ps.ByName("messageId")
    messageId, err := strconv.ParseInt(messageIdStr, 10, 64)
    if err != nil {
        http.Error(w, "Invalid message ID", http.StatusBadRequest)
        return
    }
	
    err = rt.db.Removemessage(uint64(messageId))
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
		return
        }

    w.WriteHeader(http.StatusNoContent)
}////204 no content 404, i need mess id, add auth

func (rt *_router) forwardMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	messageIdStr := ps.ByName("messageId")
	messageId, err := strconv.ParseUint(messageIdStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid message ID", http.StatusBadRequest)
		return
	}
	var requestBody struct {
		TargetChatId uint64 `json:"targeChatId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	if requestBody.TargetChatId == 0 {
		http.Error(w, "Target user ID must be provided", http.StatusUnprocessableEntity)
		return
	}

	dbmessage, err := rt.db.GetMessageById(messageId)
	if err != nil {
		http.Error(w, "Message not found", http.StatusNotFound)
		return
	}

	var message Message
	currentTime := time.Now()
	message.MessageDate = currentTime.Format("2006-01-02")
	message.MessageTime = currentTime.Format("15:04")
	message.State = "send"

	dbmessage, err = rt.db.Sendmessage(message.MessageToDatabase())
	if err != nil {
		http.Error(w, "Error forwarding message", http.StatusInternalServerError)
		return
	}
	message.MessageFromDatabase(dbmessage)
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(message)
}
//201 404 


func (rt *_router) commentMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")
	var comment Comment

	
	messageIdStr := ps.ByName("messageId")
	messageId, err := strconv.ParseUint(messageIdStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid message ID", http.StatusBadRequest)
		return
	}
	_, err = rt.db.GetMessageById(messageId)
	if err != nil {
		http.Error(w, "Message not found", http.StatusNotFound)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&comment)
	if err != nil || comment.Content == "" || comment.MessageId == 0 {
		http.Error(w, "Invalid or empty comment content", http.StatusBadRequest)
		return
	}
	
	dbcomment, err := rt.db.Commentmessage(comment.CommentToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	comment.CommentFromDatabase(dbcomment)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(comment)

}///201 404

func (rt *_router) uncommentMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	
	commentIdStr := ps.ByName("commentId")
    commentId, err := strconv.ParseInt(commentIdStr, 10, 64)
    if err != nil {
        http.Error(w, "Invalid comment ID", http.StatusBadRequest)
        return
    }
	//sprawdzic czy kom istnieje
    err = rt.db.Removecomment(uint64(commentId))
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
		return
        }

    w.WriteHeader(http.StatusNoContent)
}//204 404


func (rt *_router) getConversation(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")
	chatIdStr := ps.ByName("chatId")
	chatId, err := strconv.ParseUint(chatIdStr, 10, 64)

	conversation, err := rt.db.GetConversation(chatId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching conversation: %v", err), http.StatusInternalServerError)
		return
	}
   

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(conversation)
}