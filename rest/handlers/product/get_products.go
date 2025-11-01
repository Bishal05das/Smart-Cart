package product

import (
	"fmt"
	"net/http"

	"github.com/bishal05das/ecommerce-project/util"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	productList, err := h.productRepo.List()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	util.SendData(w,productList, 200)

}

