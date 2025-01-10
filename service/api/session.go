package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"encoding/json"

)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	w.Header().Set("content-Type", "application/json")
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dbuser, err := rt.db.CreateLogin(user.ToDatabase())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.FromDatabase(dbuser)
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(user)
}