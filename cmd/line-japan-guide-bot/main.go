package main

import (
	"flag"
	"log"

	"github.com/keijiyoshida/line-japan-guide-bot/bot"
	"github.com/keijiyoshida/line-japan-guide-bot/config"
	"github.com/keijiyoshida/line-japan-guide-bot/server"
)

// confPath retrieves a configuration file path from
// the command-line flag and returns it.
func confPath() string {
	path := flag.String("c", "", "configuration file path")
	flag.Parse()
	return *path
}

func main() {
	conf, err := config.New(confPath())
	if err != nil {
		log.Fatalln(err)
	}

	b, err := bot.New(conf.Bot)
	if err != nil {
		log.Fatalln(err)
	}

	if err := server.Run(conf.HTTP, b); err != nil {
		log.Fatalln(err)
	}
}
