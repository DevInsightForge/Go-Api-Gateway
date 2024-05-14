package mapping

import (
	"log"
	"net/http"

	"api-gateway/internal/config"
	"api-gateway/internal/proxy"
)

func RegisterMappings(mux *http.ServeMux, cfg *config.Config) error {
	for _, route := range cfg.Mappings {
		p, err := proxy.NewProxy(route)
		if err != nil {
			log.Printf("[Mapping] Error creating proxy for route %s: %v", route.Prefix, err)
			continue
		}
		mux.Handle("/"+route.Prefix+"/", p)
		log.Printf("[Mapping] Registered route prefix: '%s' to target server: %s", route.Prefix, route.Server)

	}
	return nil
}
