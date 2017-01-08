package server

import (
	"log"
	"net/http"
)

// Run runs an HTTP server.
func Run(addr, certFile, keyFile string, handler http.Handler) error {
	http.Handle("/", handler)
	log.Printf("About to listen on %s", addr)
	return http.ListenAndServeTLS(addr, certFile, keyFile, nil)
}
