package api


import (
	"github.com/julienschmidt/httprouter"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"net/http"
	"encoding/json"
	"strconv"

)


func (rt *_router) addToGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
    w.Header().Set("Content-Type", "application/json")

   
    chatIdStr := ps.ByName("chatId")
    chatId, err := strconv.ParseUint(chatIdStr, 10, 64)
    if err != nil {
        http.Error(w, "Invalid chat ID", http.StatusBadRequest)
        return
    }

   
    var requestBody struct {
        UserId uint64 `json:"userId"`
    }

    err = json.NewDecoder(r.Body).Decode(&requestBody)
    if err != nil || requestBody.UserId == 0 {
        http.Error(w, "Invalid request body or missing userId", http.StatusBadRequest)
        return
    }


    err = rt.db.AddUserToChat(chatId, requestBody.UserId)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            http.Error(w, "Chat or user not found", http.StatusNotFound)
            return
        }
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

 
    w.WriteHeader(http.StatusCreated)
    _, _ = w.Write([]byte(`{"message": "User added to chat successfully"}`))
}
