package commands

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/lixiangzhong/dnsutil"
	tele "gopkg.in/telebot.v3"
	"io"
	"os"
	"strings"
)

func Dig(c tele.Context) error {
	fields := strings.Fields(c.Message().Text)
	digCmd := flag.NewFlagSet("dig", flag.ContinueOnError)
	digType := digCmd.String("type", "A", "Query record requested [A,AAAA,NS,MX,TXT,PTR,CNAME]")
	digServer := digCmd.String("server", "8.8.8.8", "DNS Server to send the request")

	if len(fields) < 2 {
		return c.Reply("Usage: /dig \\<parameters\\> query\n\nUse `/dig \\-\\-help` to list the parameters", tele.ModeMarkdownV2)
	}
	query := fields[len(fields)-1]
	// flag processing prints to stderr on errors or help, so stderr must be captured, dump into a variable and restored to show it to the user
	old := os.Stderr // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stderr = w
	err := digCmd.Parse(fields[1:])
	outC := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	// back to normal state
	w.Close()
	os.Stderr = old // restoring the real stdout
	val := <-outC
	if err != nil {
		return c.Reply(val)
	}
	var dig dnsutil.Dig
	dig.SetDNS(*digServer)
	var result interface{}
	err = nil
	switch *digType {
	case "A":
		result, err = dig.A(query)
	case "AAAA":
		result, err = dig.AAAA(query)
	case "NS":
		result, err = dig.NS(query)
	case "MX":
		result, err = dig.MX(query)
	case "TXT":
		result, err = dig.TXT(query)
	case "PTR":
		result, err = dig.PTR(query)
	case "CNAME":
		result, err = dig.CNAME(query)
	default:
		return c.Reply(fmt.Sprintf("Invalid type %s", *digType))
	}
	if err != nil {
		return c.Reply("Error processing query")
	}
	ret := fmt.Sprintf("%+q\n", result)
	ret = strings.Replace(strings.Replace(strings.Replace(ret, "\" \"", "\n", -1), "[\"", "", -1), "\"]", "", -1)
	ret = strings.Replace(ret, "\\t", "\t", -1)
	return c.Reply(ret)
}
