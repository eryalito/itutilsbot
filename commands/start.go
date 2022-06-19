package commands

import (
	tele "gopkg.in/telebot.v3"
)

func Start(c tele.Context) error {
	return c.Send("Welcome! Send /help for more useful info about this bot :)")
}
