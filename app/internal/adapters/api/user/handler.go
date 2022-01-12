package user

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
	router.POST("/user/new", h.SignUp)
	router.POST("/user/entry", h.SignIn)
	router.GET("/user/exit", h.SignOut)
	router.PUT("/user/refresh", h.Refresh)
	router.GET("/user/list", jwt.Middleware(h.List))
	router.GET("/user/profile", jwt.Middleware(h.Profile))
	router.GET("/user/upgrade", jwt.Middleware(h.UpgradeRole))

}
