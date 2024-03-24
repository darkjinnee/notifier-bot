package notifierbot

import goerr "github.com/darkjinnee/go-err"

type Chat struct {
	Number   int64  `json:"number"`
	Status   string `json:"status"`
	Username string `json:"username"`
}

var chats = make(map[int64]Chat)
var statuses = map[int64]string{
	0: "new",
	1: "job",
}

func NewChat(chatId int64, username string) {
	if _, ok := chats[chatId]; !ok {
		chat := Chat{
			Number:   chatId,
			Username: username,
			Status:   statuses[0],
		}

		err := CreateChat(chat)
		goerr.Log(
			err,
			"[Error] notifierbot.NewChat: Failed create chat",
		)

		if err == nil {
			chats[chatId] = chat
		}
	}
}
