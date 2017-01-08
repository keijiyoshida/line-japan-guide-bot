package main

import (
	"log"

	"github.com/keijiyoshida/line-japan-guide-bot/bot"
	"github.com/keijiyoshida/line-japan-guide-bot/server"
)

func main() {
	if err := server.Run(":80", "", "", bot.New()); err != nil {
		log.Fatalln(err)
	}
}
