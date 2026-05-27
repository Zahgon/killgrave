package http

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	killgrave "github.com/friendsofgo/killgrave/internal"
)

// Proxy represent reverse proxy server.
type Proxy struct {
	server *httputil.ReverseProxy
	mode   killgrave.ProxyMode
	url    *url.URL
}

// NewProxy creates new proxy server.
func NewProxy(rawurl string, mode killgrave.ProxyMode) (*Proxy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Handler returns handler that sends request to another server.
func (p *Proxy) Handler() http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}
