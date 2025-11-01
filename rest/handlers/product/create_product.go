package product

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bishal05das/ecommerce-project/repo"
	"github.com/bishal05das/ecommerce-project/util"
)

type ReqCreateProduct struct {
	ID 	   		int		`json:"id"`
	Title  		string	`json:"title"`
	Description string	`json:"description"`
	Price 		float64 `json:"price"`
	ImgUrl 	    string  `json:"imageUrl"`
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct ReqCreateProduct
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "plz give me valid json", 400)
		return
	}

	createdProduct, err := h.productRepo.Create(repo.Product{
		Title:       newProduct.Title,
		Description: newProduct.Description,
		Price:       newProduct.Price,
		ImgUrl:      newProduct.ImgUrl,
	}) //database.Store(newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "internal server error", 500)
		return
	}

	util.SendData(w, createdProduct, http.StatusCreated)

}
