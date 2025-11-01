package product

import (
	"net/http"
	"strconv"

	"github.com/bishal05das/ecommerce-project/util"
)

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	id, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "please give me a valid product id", http.StatusBadRequest)
		return
	}

	product,err := h.productRepo.Get(id)//database.Get(id)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	if product != nil {
		util.SendData(w, product, 200)
		return
	}
	util.SendData(w, "data pai nai", http.StatusNotFound)

}
