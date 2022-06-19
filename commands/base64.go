package commands

import(
	"bytes"
	"flag"
	"io"
	"os"
	"strings"
	
    b64 "encoding/base64"
	tele "gopkg.in/telebot.v3"
)

func Base64(c tele.Context) error {
	fields := strings.Fields(c.Message().Text)
	digCmd := flag.NewFlagSet("base64", flag.ContinueOnError)
	decodeFlag := digCmd.Bool("d", false, "Set in decode mode")
	if len(fields) < 2 {
		return c.Reply("Usage: /base64 \\<parameters\\> text\n\nUse `/base64 \\-\\-help` to list the parameters", tele.ModeMarkdownV2 )
	}
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
	if err!=nil {	
		return c.Reply(val)
	}
	
	var data string
	if *decodeFlag {
		data = strings.Join(fields[2:], "")
	} else {
		data = c.Message().Text[len("/base64 "):]
	}
	if *decodeFlag {
		bytes, err := b64.StdEncoding.DecodeString(data)
		if err != nil {
			return c.Reply("Invalid base64 string")
		}
		data = string(bytes)
	} else {
		data = b64.StdEncoding.EncodeToString([]byte(data))
	}
	return c.Reply(data)
}
