package service

import (
	"fmt"
	"net/http"

	"github.com/VishwasMallikarjuna/UrlShortener/model"
	"github.com/gorilla/mux"
)

// GetOriginalURL writes original link to the body of end point
func GetOriginalURL(nus *model.UrlShortener) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("---Running in getOriginalURL---")
		vars := mux.Vars(r)
		short, ok := vars["short"]
		if !ok {
			fmt.Println("Could not retrive the short URL: ", short)
			return
		}
		w.Write([]byte(nus.UrlMap[short]))
	}
}
