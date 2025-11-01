package product

import (
	"net/http"
	"strconv"

	"github.com/bishal05das/ecommerce-project/util"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pID, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "please give me valid product id", 400)
		return
	}
	err = h.productRepo.Delete(pID)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	util.SendData(w, "successfully deleted product", 200)

}
