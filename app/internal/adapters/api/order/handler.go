package order

import (
	"site/internal/adapters/api"
	"site/pkg/middleware/jwt"

	"github.com/julienschmidt/httprouter"
)

type handler struct {
	orderService Service
	JWTHelper    jwt.Helper
}

func NewHandler(service Service, helper jwt.Helper) api.Handler {
	return &handler{orderService: service, JWTHelper: helper}
}

func (h *handler) Register(router *httprouter.Router) {
	router.POST("/order/create", h.Create)
	router.PUT("/order/update", h.Update)
	router.GET("/order/get", h.Get)
}
