package user

import (
	"github.com/bishal05das/ecommerce-project/config"
	"github.com/bishal05das/ecommerce-project/repo"
)

type Handler struct {
	cnf      *config.Config
	userRepo repo.UserRepo
}

func NewHandler(cnf *config.Config,userRepo repo.UserRepo) *Handler {
	return &Handler{
		cnf:      cnf,
		userRepo: userRepo,
	}
}
