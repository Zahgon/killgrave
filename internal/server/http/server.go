package http

import (
	_ "embed"
	"net/http"

	killgrave "github.com/friendsofgo/killgrave/internal"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//go:embed cert/server.key
var serverKey []byte

//go:embed cert/server.cert
var serverCert []byte

var (
	defaultCORSMethods        = []string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE", "PATCH", "TRACE", "CONNECT"}
	defaultCORSHeaders        = []string{"X-Requested-With", "Content-Type", "Authorization"}
	defaultCORSExposedHeaders = []string{"Cache-Control", "Content-Language", "Content-Type", "Expires", "Last-Modified", "Pragma"}
)

// ServerOpt function that allow modify the current server
type ServerOpt func(s *Server)

// Server definition of mock server
type Server struct {
	router     *mux.Router
	httpServer *http.Server
	proxy      *Proxy
	secure     bool
	imposterFs ImposterFs
}

// NewServer initialize the mock server
func NewServer(r *mux.Router, httpServer *http.Server, proxyServer *Proxy, secure bool, fs ImposterFs) Server {
	_ = "STUB: not implemented"
	return *new(Server)
}

// PrepareAccessControl Return options to initialize the mock server with default access control
func PrepareAccessControl(config killgrave.ConfigCORS) (h []handlers.CORSOption) {
	_ = "STUB: not implemented"
	return nil
}

// Build read all the files on the impostersPath and add different
// handlers for each imposter
func (s *Server) Build() error { _ = "STUB: not implemented"; return nil }

// not necessary load the imposters if you will use the tool as a proxy

// Run launch a previous configured http server if any error happens while the starting process
// application will be crashed
func (s *Server) Run() { _ = "STUB: not implemented"; return }

func (s *Server) run(secure bool) error { _ = "STUB: not implemented"; return nil }

// Shutdown shutdowns the current http server
func (s *Server) Shutdown() error { _ = "STUB: not implemented"; return nil }

func (s *Server) addImposterHandler(imposters []Imposter) { _ = "STUB: not implemented"; return }

func (s *Server) handleAll(h http.HandlerFunc) { _ = "STUB: not implemented"; return }
