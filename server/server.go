package server

import (
	"net/http"

	"github.com/keijiyoshida/line-japan-guide-bot/config"
)

// Run runs an HTTP server.
func Run(conf config.HTTP, handler http.Handler) error {
	http.Handle("/", handler)
	return http.ListenAndServeTLS(conf.Addr, conf.CertFile, conf.KeyFile, nil)
}
