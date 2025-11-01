package middleware

import (
	"fmt"
	"net/http"
)

func Hudai(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){
		fmt.Println("hudai middleware print korlam")
		next.ServeHTTP(w, r)
	})
}
