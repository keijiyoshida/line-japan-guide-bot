package worker

import (
	"sync"

	"github.com/keijiyoshida/line-japan-guide-bot/config"
	"github.com/keijiyoshida/line-japan-guide-bot/log"
	"github.com/line/line-bot-sdk-go/linebot"
)

// Log types
var (
	logTypeEvent = "event"
)

// Worker represents a worker.
type Worker struct {
	cli    *linebot.Client
	wg     *sync.WaitGroup
	evchan <-chan *linebot.Event
}

// Run runs the worker.
func (w *Worker) Run() {
	defer w.wg.Done()
	for ev := range w.evchan {
		log.JSON(logTypeEvent, ev)
	}
}

// New creates a worker and returns it.
func New(conf config.LINEClient, wg *sync.WaitGroup, evchan <-chan *linebot.Event) (*Worker, error) {
	cli, err := linebot.New(conf.ChannelSecret, conf.ChannelToken)
	if err != nil {
		return nil, err
	}
	return &Worker{cli, wg, evchan}, nil
}
