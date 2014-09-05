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
	handler := CacheHandler(fsH, cfg.CacheTime)
	if cfg.CORS {
		handler = CORSHandler(handler)
	}
	handler = LogHandler(handler)
	err := http.ListenAndServe(cfg.Address, handler)
	log.Fatal(err)
}