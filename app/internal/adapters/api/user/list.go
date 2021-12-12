package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"site/app/settings"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// List
func (h *handler) List(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	part, err := strconv.Atoi(r.FormValue("part"))

	if err != nil {
		log.Println(err.Error())
		fmt.Println("")
		http.Error(w, "Internal Server Error", 500)
	}

	fmt.Println("POST user-list: part ", part)

	offset := settings.USER_ROW_COUNT * part

	users, err := h.userService.GetAll(offset, settings.USER_ROW_COUNT)

	if err != nil {
		log.Println(err.Error())
		fmt.Println("")
		http.Error(w, "Internal Server Error", 500)
	}

	json.NewEncoder(w).Encode(users)
}
