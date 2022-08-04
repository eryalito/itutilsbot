package commands

import (
	tele "gopkg.in/telebot.v3"
)

func Help(c tele.Context) error {
	return c.Send("This is a bot made up to provide basic IT utilities like dns queries, man search and more")
}
