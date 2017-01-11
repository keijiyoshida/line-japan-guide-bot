package config

// HTTPServer represents configuration of an HTTP httpserver.
type HTTPServer struct {
	Addr     string
	CertFile string
	KeyFile  string
}
