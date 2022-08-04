package commands

import (
	"fmt"
	"strings"

	tele "gopkg.in/telebot.v3"
)

func RFC(c tele.Context) error {
	fields := strings.Fields(c.Message().Text)
	if len(fields) < 2 {
		return c.Reply("Usage: /rfc <rfc#>\n\nPowered by https://www.ietf.org/rfc")
	}
	url := fmt.Sprintf("https://www.ietf.org/rfc/rfc%s.txt.pdf", fields[1])
	a := &tele.Document{File: tele.FromURL(url), Caption: url}
	err := c.Reply(a)
	if err != nil {
		return c.Reply("Unable to find the RFC")
	}
	return nil
}
