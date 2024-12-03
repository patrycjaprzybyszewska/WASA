
import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"encoding/json"
)

type Mesage struct{
	MessageId int
	UserId int
	Content string
	MessageDate string
}


