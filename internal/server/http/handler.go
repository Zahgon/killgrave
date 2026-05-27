package http

import (
	"net/http"
)

// ImposterHandler create specific handler for the received imposter
func ImposterHandler(i Imposter) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

func writeHeaders(r Response, w http.ResponseWriter) { _ = "STUB: not implemented"; return }

func writeBody(i Imposter, r Response, w http.ResponseWriter) { _ = "STUB: not implemented"; return }

func fetchBodyFromFile(bodyFile string) (bytes []byte) { _ = "STUB: not implemented"; return nil }
