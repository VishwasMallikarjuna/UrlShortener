package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/VishwasMallikarjuna/UrlShortener/model"
)

// GetShortenedURL writes shortend link to the body of end point
func GetShortenedURL(nus *model.UrlShortener, m map[string]int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("---Running in GetShortenedURL---")
		var reqBody model.Request
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			fmt.Println("Error in parsing request body: ", err)
			return
		}

		fmt.Println("url input: ", reqBody.Url)
		if _, ok := nus.CheckMap[reqBody.Url]; ok {
			w.Write([]byte("Provided url exists. Shorted URL is: " + nus.CheckMap[reqBody.Url]))
			return
		}
		countDomine(reqBody.Url, m)
		short := RandStr(5)

		nus.Mutex.Lock()
		nus.UrlMap[short] = reqBody.Url
		nus.CheckMap[reqBody.Url] = short
		nus.Mutex.Unlock()
		w.Write([]byte("http://localhost:8192/original/" + short))

	}
}
