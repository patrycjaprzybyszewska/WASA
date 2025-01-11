package api


import (
	"github.com/julienschmidt/httprouter"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"net/http"
	"encoding/json"
	"strconv"
	"time"
)

func (rt *_router) sendMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")
	var message Message

	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if message.Content == "" || message.UserId == 0 {
		http.Error(w, "Missing required fields: content or userId", http.StatusBadRequest)
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
