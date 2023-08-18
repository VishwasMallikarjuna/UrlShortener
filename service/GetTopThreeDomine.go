package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sort"
)

// GetTop3DomineNames taks map of count of domine name and writes to body of API
// top 3 domine name and their count
func GetTop3DomineNames(domineMap map[string]int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("---Fetching Top 3 Domine names---")
		top3 := getTop3(domineMap)
		jsonv, _ := json.Marshal(top3)
		w.Write(jsonv)
	}
}

// countDomine Every unique request count is updated
func countDomine(urls string, m map[string]int) {
	parsedUrl, err := url.Parse(urls)
	if err != nil {
		log.Fatal(err)
	}
	_, found := m[parsedUrl.Hostname()]
	if found {
		m[parsedUrl.Hostname()] += 1
	} else {
		m[parsedUrl.Hostname()] = 1
	}

}

// getTop3 takes map as input and returns top 3 domines shortend
func getTop3(domineMap map[string]int) map[string]int {
	resultMap := make(map[string]int)

	if len(domineMap) > 3 {
		countlist := []int{}
		for _, count := range domineMap {
			countlist = append(countlist, count)
		}
		sort.Slice(countlist, func(i, j int) bool {
			return countlist[i] > countlist[j]
		})
		for i := 0; i < 3; i++ {
			for domineName, count := range domineMap {
				if count == countlist[i] {
					resultMap[domineName] = count
				}
			}
		}
		return resultMap
	} else {
		return domineMap
	}

}
