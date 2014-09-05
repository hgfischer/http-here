package main

import "net/http"

type corsHandler struct {
	handler http.Handler
}

func newCORSHandler(h http.Handler) http.Handler {
	return &corsHandler{h}
}

func (h *corsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	h.handler.ServeHTTP(w, r)
}
