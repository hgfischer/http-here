package main

import (
	"fmt"
	"net/http"
)

type cacheHandler struct {
	handler   http.Handler
	cacheTime int
}

func newCacheHandler(h http.Handler, cacheTime int) http.Handler {
	return &cacheHandler{h, cacheTime}
}

func (h *cacheHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.cacheTime >= 0 {
		w.Header().Set("Cache-Control", fmt.Sprintf("max-age=%d", h.cacheTime))
	} else {
		w.Header().Set("Cache-Control", "private,max-age=0,no-cache")
	}
	h.handler.ServeHTTP(w, r)
}
