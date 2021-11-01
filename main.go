package main

import (
	"fmt"
	"os"

	"github.com/gempir/go-twitch-irc"
)

// use environment variables to access data

func main() {
	// init environment variables
	botName := os.Getenv("BOT_USERNAME")
	botOAuthToken := os.Getenv("BOT_TOKEN")
	channelName := os.Getenv("CHANNEL_NAME")

	client := twitch.NewClient(botName, fmt.Sprintf("oauth: %s", botOAuthToken))
	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		fmt.Println(message.Message)
	})

	client.Join(channelName)

	err := client.Connect()
	if err != nil {
		panic(err)
	}
}
