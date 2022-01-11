package message

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type listParams struct {
	Username string `json:"username" example:"example" valid:"type(string)"`
	Password string `json:"password" example:"123456789" valid:"type(string)"`
}

func (h *handler) List(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
