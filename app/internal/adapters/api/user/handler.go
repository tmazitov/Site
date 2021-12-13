package user

import (
	"site/app/internal/adapters/api"
	"site/app/pkg/middleware/jwt"

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
	router.POST("/user/new", h.SignUp)
	router.POST("/user/entry", h.SignIn)
	router.POST("/user/list", h.List)
}
