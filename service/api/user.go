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
	authHeader := r.Header.Get("Authorization")
    if authHeader == "" {
        http.Error(w, "Missing authorization", http.StatusUnauthorized)
        return
    }
	authid, err := auth(authHeader)
    if err != nil {
        http.Error(w, "Invalid token", http.StatusUnauthorized)
        return
	}
	if authid != userId {
		http.Error(w, "bad autorization", http.StatusUnauthorized)
		return
	}
	var user User
	user.UserId = userId
	userPhoto, err := rt.db.GetUserPhotoById(userId)
    user.UserPhoto = userPhoto
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}	
	dbuser, err := rt.db.SetUsername(user.ToDatabase(), user.UserName)
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
	authHeader := r.Header.Get("Authorization")
    if authHeader == "" {
        http.Error(w, "Missing authorization", http.StatusUnauthorized)
        return
    }
	authid, err := auth(authHeader)
    if err != nil {
        http.Error(w, "Invalid token", http.StatusUnauthorized)
        return
	}
	if authid != userId {
		http.Error(w, "bad autorization", http.StatusUnauthorized)
		return
	}
	user.UserId = userId
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}	
	userName, err := rt.db.GetUserNameById(userId)
    if err != nil {
        http.Error(w, "No name", http.StatusInternalServerError)
        return
    }
	user.UserName = userName

	dbuser, err := rt.db.SetUserphoto(user.ToDatabase(), user.UserPhoto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.FromDatabase(dbuser)
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(user)
}
