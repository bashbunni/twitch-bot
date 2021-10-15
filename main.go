package main

import (
	"fmt"

	"github.com/gempir/go-twitch-irc"
)

func main() {
	client := twitch.NewClient("zshbunni", fmt.Sprintf("oauth: %s", oauthToken))
	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		fmt.Println(message.Message)
	})

	client.Join("bashbunni") // join bb channel

	err := client.Connect()
	if err != nil {
		panic(err)
	}
}
