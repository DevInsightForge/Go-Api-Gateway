package app

import (
	"fmt"
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

	appCfg := config.GetAppConfig()
	fullAddr := fmt.Sprintf("%s:%s", appCfg.ServerAddr, appCfg.ServerPort)

	httpServer := &http.Server{
		Addr:    fullAddr,
		Handler: handler,
	}

	// Run the server
	log.Printf("[App] Server is running at http://%s\n", fullAddr)
	err = httpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
