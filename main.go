package main

import (
	"log"
	"os"
	"time"
	"github.com/eryalus/itutilsbot/commands"

	tele "gopkg.in/telebot.v3"
)

func main() {
	API_TOKEN, ok := os.LookupEnv("API_TOKEN")
	if !ok {
		log.Fatal("API_TOKEN not provided")
	}

	pref := tele.Settings{
		Token:  API_TOKEN,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}
	
	b.Handle("/start", commands.Start)
	b.Handle("/help", commands.Help)
	b.Handle("/dig", commands.Dig)

	b.Start()
}