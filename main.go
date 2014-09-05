package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	cfg := loadConfig()
	fmt.Println("Starting up http-here at", cfg.Address)
	fmt.Println("Hit CTRL-C to stop this server")

	fsH := http.FileServer(http.Dir(""))
	handler := newCacheHandler(fsH, cfg.CacheTime)
	if cfg.CORS {
		handler = newCORSHandler(handler)
	}
	handler = newLogHandler(handler)
	err := http.ListenAndServe(cfg.Address, handler)
	log.Fatal(err)
}
