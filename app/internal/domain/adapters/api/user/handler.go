package user

import (
	"site/app/internal/domain/adapters/api"

	"github.com/julienschmidt/httprouter"
)

type handler struct {
	userService Service
}

func NewHandler(service Service) api.Handler {
	return &handler{userService: service}
}

func (h *handler) Register(router *httprouter.Router) {
	router.POST("/user/new", h.SignUp)
	router.POST("/user/entry", h.SignIn)
	router.POST("/user/list", h.List)
}
