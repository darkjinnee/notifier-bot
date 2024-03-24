package tgbot

import (
	goerr "github.com/darkjinnee/go-err"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func New(token string, debug bool) *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(token)
	goerr.Fatal(
		err,
		"[Error] tgbot.New: Failed to create bot",
	)

	bot.Debug = debug
	log.Printf(
		"[X] tgbot.New: Authorized on account: %s",
		bot.Self.UserName,
	)

	return bot
}

func GetUpdateConf(timeout int) tgbotapi.UpdateConfig {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = timeout

	return u
}
