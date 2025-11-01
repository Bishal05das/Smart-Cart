package product

import (
	"net/http"

	middleware "github.com/bishal05das/ecommerce-project/rest/middlewares"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /products", manager.With(http.HandlerFunc(h.GetProducts)))
	mux.Handle("POST /products", manager.With(http.HandlerFunc(h.CreateProduct), h.middleware.Authentication))
	mux.Handle("GET /products/{id}", manager.With(http.HandlerFunc(h.GetProduct)))
	mux.Handle("PUT /products/{id}", manager.With(http.HandlerFunc(h.UpdateProduct)))
	mux.Handle("DELETE /products/{id}", manager.With(http.HandlerFunc(h.DeleteProduct), h.middleware.Authentication))

}
