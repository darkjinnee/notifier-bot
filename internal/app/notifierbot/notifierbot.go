package notifierbot

import (
	"fmt"
	"strings"
)

func Run() {
	fmt.Printf(
		"Hey! I am a telegram bot \"%s\", for sending notifications\n",
		strings.ToTitle(Conf.App.Name),
	)

	h := []Header{
		{
			Key:   "Accept",
			Value: "application/json",
		},
		{
			Key:   "Content-Type",
			Value: "application/json",
		},
	}

	r := []Route{
		{
			Headers: h,
			Method:  "GET",
			Pattern: "/",
			Handler: Home,
		},
	}
	Listen(r)
}
