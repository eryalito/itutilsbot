package main

import (
	"log"
	"os"
	"time"

	"github.com/eryalus/itutilsbot/commands"
	"github.com/eryalus/itutilsbot/policies"

	tele "gopkg.in/telebot.v3"
)

func main() {

	var auth_config policies.Authconfig
	auth_config.GetConfg()
	auth_config.Init()

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
	middleware := policies.GetMiddlewareFunc_Allow(auth_config)
	b.Handle("/start", commands.Start, middleware)
	b.Handle("/help", commands.Help, middleware)
	b.Handle("/base64", commands.Base64, middleware)
	b.Handle("/dig", commands.Dig, middleware)
	b.Handle("/man", commands.Man, middleware)
	b.Handle("/rfc", commands.RFC, middleware)

	b.Start()
}
