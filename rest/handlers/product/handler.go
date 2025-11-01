package product

import (
	"github.com/bishal05das/ecommerce-project/repo"
	middleware "github.com/bishal05das/ecommerce-project/rest/middlewares"
)

type Handler struct {
	middleware  *middleware.Middlewares
	productRepo repo.ProductRepo
}

func NewHandler(middleware *middleware.Middlewares,productRepo repo.ProductRepo) *Handler {
	return &Handler{
		middleware:  middleware,
		productRepo: productRepo,
	}
}
