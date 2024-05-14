package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

type Proxy struct {
	Target *url.URL
	Proxy  *httputil.ReverseProxy
	Prefix string
}

func NewProxy(target, prefix string) (*Proxy, error) {
	targetURL, err := url.Parse(target)
	if err != nil {
		return nil, err
	}

	return &Proxy{
		Target: targetURL,
		Proxy:  httputil.NewSingleHostReverseProxy(targetURL),
		Prefix: prefix,
	}, nil
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, p.Prefix) {
		r.URL.Path = strings.TrimPrefix(r.URL.Path, p.Prefix)
	} else if p.Prefix == "/" {
		// Special case for handling the root path "/"
		r.URL.Path = strings.TrimPrefix(r.URL.Path, p.Prefix)
	} else {
		http.NotFound(w, r)
		return
	}

	r.Host = p.Target.Host

	p.Proxy.ServeHTTP(w, r)
}
