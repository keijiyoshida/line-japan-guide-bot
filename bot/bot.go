package bot

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/keijiyoshida/line-japan-guide-bot/config"
	"github.com/line/line-bot-sdk-go/linebot"
)

// Bot represents a LINE bot interface.
type Bot interface {
	http.Handler
}

// bot represents a LINE bot.
type bot struct {
	cli *linebot.Client
}

func (b *bot) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	events, err := b.cli.ParseRequest(r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		log.Printf("error %q\n", err)
		return
	}

	for _, event := range events {
		d, err := json.Marshal(event)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("error %q\n", err)
			return
		}

		log.Println("event", string(d))

		switch event.Type {
		case linebot.EventTypeFollow:
			_, err := b.cli.ReplyMessage(
				event.ReplyToken,
				linebot.NewTextMessage("Welcome to Japan and thank you for following Japan Guide!"),
			).Do()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Printf("error %q\n", err)
				return
			}
		case linebot.EventTypeUnfollow:
		case linebot.EventTypeJoin:
		case linebot.EventTypeLeave:
		case linebot.EventTypeMessage:
			_, err := b.cli.ReplyMessage(
				event.ReplyToken,
				linebot.NewTextMessage("Hi!"),
			).Do()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Printf("error %q\n", err)
				return
			}
		}
	}
}

func (b *bot) serve(w http.ResponseWriter, r *http.Request) (int, error) {
	events, err := b.cli.ParseRequest(r)
	switch {
	case err == linebot.ErrInvalidSignature:
		return http.StatusBadRequest, err
	case err != nil:
		return http.StatusInternalServerError, err
	}

	log.Println(events)

	return http.StatusOK, nil
}

// New creates a LINE bot and returns it.
func New(conf config.Bot) (Bot, error) {
	cli, err := linebot.New(conf.ChannelSecret, conf.ChannelToken)
	if err != nil {
		return nil, err
	}
	return &bot{cli}, nil
}
