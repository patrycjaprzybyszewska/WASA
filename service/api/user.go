package api


import (
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"encoding/json"
	"strconv"
)

func (rt *_router) setMyUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")

	userId, err := strconv.ParseUint(ps.ByName("userId"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var user User
	user.UserId = userId
	var requestBody struct {
		Name string `json:"name"`
	}
	
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}	

	if len(requestBody.Name) < 3 || len(requestBody.Name) > 16 {
		http.Error(w, "Username must be between 3 and 16 characters", http.StatusUnprocessableEntity)
		return
	}


///dodac autoryzacje
	dbuser, err := rt.db.SetUsername(user.ToDatabase(), requestBody.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.FromDatabase(dbuser)
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(user)
}

///to set username i need to get id, username change it into database, response 201



func (rt *_router) setMyPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "application/json")

	userId, err := strconv.ParseUint(ps.ByName("userId"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var user User
	user.UserId = userId
	var requestBody struct {
		Name string `json:"name"`
	}
	
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}	



///dodac autoryzacje
	dbuser, err := rt.db.SetUserphoto(user.ToDatabase(), requestBody.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.FromDatabase(dbuser)
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(user)
}
