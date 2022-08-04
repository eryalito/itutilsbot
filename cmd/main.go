package main

import (
	"github.com/eryalus/itutilsbot/internal/commands"
	"github.com/eryalus/itutilsbot/internal/config"
	"github.com/eryalus/itutilsbot/internal/policies"

	"log"
	"time"

	tele "gopkg.in/telebot.v3"
)

func main() {
	var flags config.ArgFlags
	flags.Init()
	flags.Parse()

	var middleware tele.MiddlewareFunc

	var auth_config policies.Authconfig
	if flags.Config.DisableAuth {
		middleware = policies.GetMiddlewareFunc_Bypass()
	} else {
		auth_config.GetConfg(flags.Config.AuthPath)
		auth_config.Init()
		middleware = policies.GetMiddlewareFunc_Allow(auth_config)
	}

	pref := tele.Settings{
		Token:  flags.Config.Token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}
	b.Handle("/start", commands.Start, middleware)
	b.Handle("/help", commands.Help, middleware)
	b.Handle("/base64", commands.Base64, middleware)
	b.Handle("/dig", commands.Dig, middleware)
	b.Handle("/man", commands.Man, middleware)
	b.Handle("/rfc", commands.RFC, middleware)

	b.Start()
}
