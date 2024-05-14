package app

import (
	"log"
	"net/http"

	"api-gateway/internal/config"
	"api-gateway/internal/health"
	"api-gateway/internal/mapping"
)

func Run() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("[App] could not load config: %v", err)
	}

	handler := http.NewServeMux()
	handler.HandleFunc("/health", health.HealthCheckHandler)

	err = mapping.RegisterMappings(handler, cfg)
	if err != nil {
		log.Fatalf("[App] could not register routes: %v", err)
	}

	fullAddr := "localhost:8080"
	httpServer := &http.Server{
		Addr:    fullAddr,
		Handler: handler,
	}
	log.Printf("[App] Server is running at http://%s\n", fullAddr)

	// Run the server
	err = httpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
