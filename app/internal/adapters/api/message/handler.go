package message

import (
	"site/internal/adapters/api"
	"site/pkg/middleware/jwt"

	"github.com/julienschmidt/httprouter"
)

type handler struct {
	userService Service
	JWTHelper   jwt.Helper
}

func NewHandler(service Service, helper jwt.Helper) api.Handler {
	return &handler{userService: service, JWTHelper: helper}
}

func (h *handler) Register(router *httprouter.Router) {
	router.POST("/message/create", jwt.Middleware(h.Create))
	router.POST("/message/change", jwt.Middleware(h.Change))
	router.DELETE("/message/delete", jwt.Middleware(h.Delete))
	router.GET("/message/list", jwt.Middleware(h.List))
}
