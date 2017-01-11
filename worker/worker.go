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

// Text messages
var (
	textMessageEat   = "eat"
	textMessageHelp  = "help"
	textMessageUsage = "usage"
)

var usage = `[usage]
* To find good places to eat close to your location, send a text of "eat" and share your location by tapping the "Share Location" button at the bottom.
* To see the usage of Japan Guide (this message), send a text of "usage" or "help".`

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
		if err := w.handleEvent(ev); err != nil {
			log.Error(err)
		}
	}
}

func (w *Worker) handleEvent(ev *linebot.Event) error {
	log.JSON(logTypeEvent, ev)

	switch ev.Type {
	case linebot.EventTypeFollow, linebot.EventTypeJoin:
		return w.greet(ev.ReplyToken)
	case linebot.EventTypeMessage:
		switch message := ev.Message.(type) {
		case *linebot.TextMessage:
			switch message.Text {
			case textMessageEat:
				return w.replyEat(ev.ReplyToken)
			case textMessageHelp, textMessageUsage:
				return w.showUsage(ev.ReplyToken)
			}
		}
		return w.handleUnknown(ev.ReplyToken)
	}

	return nil
}

func (w *Worker) greet(replyToken string) error {
	return w.replyMessage(replyToken, "Thank you for following Japan Guide and welcome to Japan!\n\n"+usage)
}

func (w *Worker) showUsage(replyToken string) error {
	return w.replyMessage(replyToken, usage)
}

func (w *Worker) replyMessage(replyToken string, message string) error {
	if _, err := w.cli.ReplyMessage(replyToken, linebot.NewTextMessage(message)).Do(); err != nil {
		return err
	}
	return nil
}

func (w *Worker) replyEat(replyToken string) error {
	return w.replyMessage(replyToken, "OK. I'll find good places close to your location to eat close to your location. "+
		`Please share your location by tapping the "Share Location" button at the bottom.`)
}

func (w *Worker) handleUnknown(replyToken string) error {
	return w.replyMessage(replyToken, "I'm sorrry, but I cannot understand your message.")
}

// New creates a worker and returns it.
func New(conf config.LINEClient, wg *sync.WaitGroup, evchan <-chan *linebot.Event) (*Worker, error) {
	cli, err := linebot.New(conf.ChannelSecret, conf.ChannelToken)
	if err != nil {
		return nil, err
	}
	return &Worker{cli, wg, evchan}, nil
}
