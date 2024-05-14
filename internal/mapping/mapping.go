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
			log.Printf("Error creating proxy for route %s: %v", route.Prefix, err)
			continue
		}
		mux.Handle("/"+route.Prefix+"/", p)
		log.Printf("Registered mapping for route /%s/", route.Prefix)
	}
	return nil
}
