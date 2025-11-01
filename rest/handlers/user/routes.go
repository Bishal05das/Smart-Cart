package user

import (
	"net/http"

	middleware "github.com/bishal05das/ecommerce-project/rest/middlewares"
)

func (h *Handler) RergisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	mux.Handle("POST /users", manager.With(http.HandlerFunc(h.CreateUser)))
	mux.Handle("POST /users/login", manager.With(http.HandlerFunc(h.Login)))
}
