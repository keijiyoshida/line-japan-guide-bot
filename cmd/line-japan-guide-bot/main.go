package main

import (
	"log"
	"net/http"

	"github.com/keijiyoshida/line-japan-guide-bot/bot"
)

func main() {
	http.Handle("/", bot.New())
	err := http.ListenAndServe(":80", nil)
	log.Fatalln(err)
}
