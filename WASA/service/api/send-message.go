import (
//	"github.com/julienschmidt/httprouter"
	"net/http"
	"encoding/josn"
	"math/rand"
	"strconv"
)


func (rt *_roter) sendMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content(type", "application/json")
	
	id:= len(Messages)

	Messages = append(Messages, message{
		Id: id,
		secret: rand.Intn(1000)

	})

	conversations = append(conversations, []Conversation{})

	json.NewEncoder(w).Encode(id)
}