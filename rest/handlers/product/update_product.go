package product

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/bishal05das/ecommerce-project/repo"
	"github.com/bishal05das/ecommerce-project/util"
)

type ReqUpdateProduct struct {
	ID 	   		int		`json:"id"`
	Title  		string	`json:"title"`
	Description string	`json:"description"`
	Price 		float64 `json:"price"`
	ImgUrl 	    string  `json:"imageUrl"`
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pID, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "please give me a valid product id", 400)
		return
	}

	var newProduct ReqUpdateProduct
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "plz give me valid json", 400)
		return
	}
	newProduct.ID = pID
	_,err = h.productRepo.Update(repo.Product{
		ID:          newProduct.ID,
		Title:       newProduct.Title,
		Description: newProduct.Description,
		Price:       newProduct.Price,
		ImgUrl:      newProduct.ImgUrl,
	})//database.Update(newProduct)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	util.SendData(w, "successfully updated product", 201)
}
