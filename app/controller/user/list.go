package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"site/app/model"
	"site/app/settings"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// List
func List(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	part, err := strconv.Atoi(r.FormValue("part"))
	if err != nil {
		log.Println(err.Error())
		fmt.Println("")
		http.Error(w, "Internal Server Error", 500)
	}

	fmt.Println("POST user-list : part ", part)

	start := settings.USER_ROW_COUNT * part
	end := settings.USER_ROW_COUNT * (part + 1)

	users := model.GetPartUsers(start, end)

	json.NewEncoder(w).Encode(users)
}
