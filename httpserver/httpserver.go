package httpserver

import (
	"net/http"

	"github.com/keijiyoshida/line-japan-guide-bot/config"
	"github.com/keijiyoshida/line-japan-guide-bot/log"
	"github.com/line/line-bot-sdk-go/linebot"
)

// HTTPServer represents an HTTP server.
type HTTPServer struct {
	addr          string
	certFile      string
	keyFile       string
	channelSecret string
	evchan        chan<- *linebot.Event
}

func (hs *HTTPServer) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	events, err := linebot.ParseRequest(hs.channelSecret, r)
	if err != nil {
		log.Error(err)
		if err == linebot.ErrInvalidSignature {
			rw.WriteHeader(http.StatusBadRequest)
		} else {
			rw.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	for _, ev := range events {
		hs.evchan <- ev
	}
}

// Run runs an HTTP server.
func (hs *HTTPServer) Run() error {
	http.Handle("/", hs)
	return http.ListenAndServeTLS(hs.addr, hs.certFile, hs.keyFile, nil)
}

// New creates an HTTP server and returns it.
func New(conf *config.Config, evchan chan<- *linebot.Event) *HTTPServer {
	return &HTTPServer{
		conf.HTTPServer.Addr,
		conf.HTTPServer.CertFile,
		conf.HTTPServer.KeyFile,
		conf.LINEClient.ChannelSecret,
		evchan,
	}
}
