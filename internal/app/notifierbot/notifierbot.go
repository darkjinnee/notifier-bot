package notifierbot

import (
	"fmt"
	"github.com/darkjinnee/notifierbot/pkg/adapter/httpx"
	"strings"
)

func Run() {
	fmt.Printf(
		"Hey! I am a telegram bot \"%s\", for sending notifications\n",
		strings.ToTitle(Conf.App.Name),
	)

	/*bot := tgbot.New(Conf.Bot.Token, Conf.Bot.Debug)
	u := tgbot.GetUpdateConf(Conf.Bot.Timeout)
	updatesBot := bot.GetUpdatesChan(u)

	go func() {
		for update := range updatesBot {
			if update.Message == nil {
				continue
			}

			NewChat(
				update.Message.Chat.ID,
				update.Message.From.UserName,
			)
		}
	}()*/

	r := []httpx.Route{
		{
			Headers: httpx.JSONHeaders,
			Method:  "GET",
			Pattern: "/home",
			Handler: Home,
		},
		{
			Headers: httpx.JSONHeaders,
			Method:  "GET",
			Pattern: "/test",
			Handler: Test,
		},
	}
	addr := Conf.Httpx.Host + ":" + Conf.Httpx.Port
	httpx.Listen(r, addr)
}
