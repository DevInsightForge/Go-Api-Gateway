package main

import (
	"log"
	"net/http"

	"api-gateway/internal/config"
	"api-gateway/internal/proxy"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	mux := http.NewServeMux()

	for _, route := range cfg.Mappings {
		p, err := proxy.NewProxy(route.Server, route.Prefix)
		if err != nil {
			log.Fatalf("could not register routes: %v", err)
		}
		mux.Handle(route.Prefix, p)
	}

	log.Println("Starting server on :8080")
	err = http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
