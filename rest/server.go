package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/bishal05das/ecommerce-project/config"
	"github.com/bishal05das/ecommerce-project/rest/handlers/product"
	"github.com/bishal05das/ecommerce-project/rest/handlers/user"
	middleware "github.com/bishal05das/ecommerce-project/rest/middlewares"
)

type Server struct {
	cnf            *config.Config
	productHandler *product.Handler
	userHandler    *user.Handler
}

func NewServer(cnf *config.Config,productHandler *product.Handler, userHandler *user.Handler) *Server {
	return &Server{
		cnf: cnf,
		productHandler: productHandler,
		userHandler:    userHandler,
	}
}

func (s *Server) Start() {
	mux := http.NewServeMux()
	manager := middleware.NewManager()
	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Hudai,
		middleware.Logger,
	)

	s.productHandler.RegisterRoutes(mux, manager)
	s.userHandler.RergisterRoutes(mux, manager)

	wrappedMux := manager.WrapMux(mux)
	addr := ":" + strconv.Itoa(s.cnf.HttpPort)
	fmt.Println("Server running on", addr)
	err := http.ListenAndServe(addr, wrappedMux)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}
