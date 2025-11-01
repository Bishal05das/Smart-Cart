package middleware

import "github.com/bishal05das/ecommerce-project/config"

type Middlewares struct {
	cnf *config.Config
}

func NewMiddlewares(cnf *config.Config) *Middlewares {
	return &Middlewares{cnf: cnf}
}