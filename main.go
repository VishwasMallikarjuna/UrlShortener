package main

import (
	"log"
	"net/http"

	"github.com/VishwasMallikarjuna/UrlShortener/model"
	"github.com/VishwasMallikarjuna/UrlShortener/service"
	"github.com/gorilla/mux"
)

func main() {
	nus := model.NewURLShortener()
	router := mux.NewRouter()
	domineMap := make(map[string]int)
	router.HandleFunc("/short", service.GetShortenedURL(nus, domineMap)).Methods("POST")
	router.HandleFunc("/original/{short}", service.GetOriginalURL(nus))
	router.HandleFunc("/top3Dns", service.GetTop3DomineNames(domineMap))
	log.Fatal(http.ListenAndServe(":8192", router))
}
