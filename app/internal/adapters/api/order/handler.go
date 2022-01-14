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
	router.POST("/order/create", jwt.Middleware(h.Create))
	router.PUT("/order/update", jwt.Middleware(h.Update))
	router.GET("/order/get", jwt.Middleware(h.Get))
	router.GET("/order/list", jwt.Middleware(h.List))
	router.DELETE("/order/delete", jwt.Middleware(h.Delete))
	router.PUT("/order/complite", jwt.Middleware(h.Complite))
}
