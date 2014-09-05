package main

import (
	"fmt"
	"net/http"
)

type cacheHandler struct {
	handler   http.Handler
	cacheTime int
}

func CacheHandler(h http.Handler, cacheTime int) http.Handler {
	return &cacheHandler{h, cacheTime}
}

func (h *cacheHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.cacheTime >= 0 {
		w.Header().Set("cache-control", fmt.Sprintf("max-age=%d", h.cacheTime))
	} else {
		w.Header().Set("cache-control", "private, max-age=0, no-cache")
	}
	h.handler.ServeHTTP(w, r)
}

//cache-control:max-age=10
//Connection:keep-alive
//content-length:6
//content-type:text/html; charset=UTF-8
//Date:Fri, 05 Sep 2014 02:14:17 GMT
//etag:"525266-6-Thu Sep 04 2014 22:11:09 GMT-0300 (BRT)"
//last-modified:Fri, 05 Sep 2014 01:11:09 GMT
//server:ecstatic-0.4.13
