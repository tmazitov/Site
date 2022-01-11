package message

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type deleteParams struct {
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
