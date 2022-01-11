package message

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type changeParams struct {
}

func (h *handler) Change(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
