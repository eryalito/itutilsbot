package commands

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"math"
	"strings"

	tele "gopkg.in/telebot.v3"
)

func Man(c tele.Context) error {
	fields := strings.Fields(c.Message().Text)
	if len(fields) < 2 {
		return c.Reply("Usage: /man <command>\n\nPowered by https://man7.org/linux/man-pages")
	}
	url := fmt.Sprintf("https://man7.org/linux/man-pages/man1/%s.1.html", fields[1])
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return c.Reply("Can't fetch this command")
	}
	defer func() { // recovers panic
		if e := recover(); e != nil {
			c.Reply("Can't fetch this command")
		}
	}()
	var b bytes.Buffer
	msg := doc.Find("html body h2").Nodes[0].FirstChild.NextSibling.Data
	msg += "\n"
	err = goquery.Render(&b, doc.Find("body > pre:nth-child(8)"))
	if err != nil {
		return c.Reply("Can't fetch this command")
	}
	msg += b.String() + "\n"
	msg += doc.Find("html body h2").Nodes[1].FirstChild.NextSibling.Data
	msg += "\n"
	err = goquery.Render(&b, doc.Find("body > pre:nth-child(10)"))
	if err != nil {
		return c.Reply("Can't fetch this command")
	}
	msg += b.String()
	footer := "\n\nMore info at " + url
	// Truncate message body to fit in telegram message size
	msg = msg[:int(math.Min(float64(4000-len(footer)), float64(len(msg))))]
	msg += footer
	return c.Send(msg, tele.ModeHTML)
}
