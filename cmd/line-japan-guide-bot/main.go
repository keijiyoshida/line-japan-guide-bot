package main

import (
	"flag"
	"log"
	"sync"

	"github.com/keijiyoshida/line-japan-guide-bot/config"
	"github.com/keijiyoshida/line-japan-guide-bot/httpserver"
	"github.com/keijiyoshida/line-japan-guide-bot/worker"
	"github.com/line/line-bot-sdk-go/linebot"
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

	evchan := make(chan *linebot.Event, conf.EvchanBufSize)

	wg := new(sync.WaitGroup)
	wg.Add(conf.NumWorker)

	for i := 0; i < conf.NumWorker; i++ {
		w, err := worker.New(conf.LINEClient, wg, evchan)
		if err != nil {
			log.Fatalln(err)
		}
		go w.Run()
	}

	if err := httpserver.New(conf, evchan).Run(); err != nil {
		log.Fatalln(err)
	}
}
