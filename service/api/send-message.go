
import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"encoding/json"
	"time"
)

type Mesage struct{
	MessageId int
	UserId int
	Content string
	MessageDate string
	State string
	Content string
}

func (rt *_router) sendMessage (w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	id:= len(Messages)
	myMessages = append(Messages, Message{
		MessageId: id,
		UserId:1,
		Conetent: "",
		MessageDate: time.Now(),
		State: send,


	})

	json.NewEncoder(w).Encode(id)
}
