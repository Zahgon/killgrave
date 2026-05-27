package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

// MatcherBySchema check if the request matching with the schema file
func MatcherBySchema(imposter Imposter) mux.MatcherFunc {
	_ = "STUB: not implemented"
	return *new(mux.MatcherFunc)
}

// TODO: inject the logger

func validateSchema(imposter Imposter, req *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}
