package proxy

import (
	"api-gateway/internal/config"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

type Proxy struct {
	Prefix string
	Target *url.URL
	Proxy  *httputil.ReverseProxy
}

func NewProxy(serverMap config.ServerMap) (*Proxy, error) {
	parsedUrl, err := url.Parse(serverMap.Server)
	if err != nil {
		return nil, err
	}

	return &Proxy{
		Prefix: serverMap.Prefix,
		Target: parsedUrl,
		Proxy:  httputil.NewSingleHostReverseProxy(parsedUrl),
	}, nil
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/"+p.Prefix)
	r.URL.Scheme = p.Target.Scheme
	r.Header.Set("X-Forwarded-Host", r.Host)
	r.Host = p.Target.Host

	p.Proxy.ServeHTTP(w, r)
}
