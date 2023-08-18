package model

import "sync"

type UrlShortener struct {
	UrlMap   map[string]string
	CheckMap map[string]string
	Mutex    sync.Mutex
}

type Request struct {
	Url string `json:"url"`
}

func NewURLShortener() *UrlShortener {
	return &UrlShortener{
		UrlMap:   make(map[string]string),
		CheckMap: make(map[string]string),
	}
}
