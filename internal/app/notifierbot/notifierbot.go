package notifierbot

import (
	"fmt"
	"strings"
)

func Run() {
	fmt.Printf(
		"Hey! I am a telegram bot \"%s\", for sending notifications\n",
		strings.ToTitle(conf.App.Name),
	)
}
