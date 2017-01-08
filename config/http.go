package config

// HTTP represents configuration of an HTTP server.
type HTTP struct {
	Addr     string
	CertFile string
	KeyFile  string
}
