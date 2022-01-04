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
	router.POST("/api/user/new", h.SignUp)
	router.POST("/api/user/entry", h.SignIn)
	router.GET("/api/user/exit", h.SignOut)
	router.PUT("/api/user/refresh", h.Refresh)
	router.GET("/api/user/list", jwt.Middleware(h.List))
	router.GET("/api/user/profile", jwt.Middleware(h.Profile))
	router.GET("/api/user/upgrade", jwt.Middleware(h.UpgradeRole))
}
