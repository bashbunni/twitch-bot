package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/joho/godotenv"
)

func getEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error loading .env file")
	}
	return os.Getenv(key)
}

func main() {
	botName := getEnvVariable("BOT_USERNAME")
	botOAuthToken := getEnvVariable("BOT_TOKEN")
	channelName := getEnvVariable("CHANNEL_NAME")

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
