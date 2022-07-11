package policies

import (
	tele "gopkg.in/telebot.v3"
)

func GetMiddlewareFunc_Allow(auth_config Authconfig) tele.MiddlewareFunc {
	return func(h tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			if msg, ok := auth_config.Validate(c.Message()); ok {
				return h(c)
			} else {
				return c.Reply(msg)
			}
		}
	}
}
