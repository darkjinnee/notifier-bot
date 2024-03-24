package notifierbot

import (
	"fmt"
	"github.com/darkjinnee/notifierbot/internal/pkg/http"
	"github.com/darkjinnee/notifierbot/internal/pkg/tgbot"
	"strings"
)

func Run() {
	fmt.Printf(
		"Hey! I am a telegram bot \"%s\", for sending notifications\n",
		strings.ToTitle(Conf.App.Name),
	)

	bot := tgbot.New(Conf.Bot.Token, Conf.Bot.Debug)
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
	}()

	h := []http.Header{
		{
			Key:   "Accept",
			Value: "application/json",
		},
		{
			Key:   "Content-Type",
			Value: "application/json",
		},
	}
	r := []http.Route{
		{
			Headers: h,
			Method:  "GET",
			Pattern: "/",
			Handler: Home,
		},
		{
			Headers: h,
			Method:  "GET",
			Pattern: "/test",
			Handler: Test,
		},
	}
	addr := Conf.Http.Host + ":" + Conf.Http.Port
	http.Listen(r, addr)
}
