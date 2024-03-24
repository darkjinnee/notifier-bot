package notifierbot

import "github.com/darkjinnee/notifierbot/internal/pkg/client"

func CreateChat(chat Chat) error {
	url := Conf.Api.URL + "/create-chat"
	_, err := client.Request(url, "POST", chat)

	return err
}
