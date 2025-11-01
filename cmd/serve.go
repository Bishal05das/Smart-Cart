package cmd

import (
	"fmt"
	"os"

	"github.com/bishal05das/ecommerce-project/config"
	"github.com/bishal05das/ecommerce-project/infra/db"
	"github.com/bishal05das/ecommerce-project/repo"
	"github.com/bishal05das/ecommerce-project/rest"
	"github.com/bishal05das/ecommerce-project/rest/handlers/product"
	"github.com/bishal05das/ecommerce-project/rest/handlers/user"
	middleware "github.com/bishal05das/ecommerce-project/rest/middlewares"
)

func Serve() {
	cnf := config.GetConfig()
	dbCon, err := db.NewConnection(cnf.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = db.MigrateDB(dbCon, cnf)
	if err != nil {
		fmt.Println("DB Migration failed:", err)
		os.Exit(1)
	}

	middleware := middleware.NewMiddlewares(cnf)

	productRepo := repo.NewProductRepo(dbCon)
	productHandler := product.NewHandler(middleware, productRepo)
	userRepo := repo.NewUserRepo(dbCon)
	userHandler := user.NewHandler(cnf, userRepo)

	server := rest.NewServer(cnf, productHandler, userHandler)
	server.Start()
}
