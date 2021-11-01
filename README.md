# Welcome to my Twitch bot
it doesn't do anything yet (lol)

## How to Use
I have ignored a file in my main package that includes the implementation of setupEnvironment. 

You're going to want to create a file for your bot containing the following:

```
func setupEnvironment() {
	os.Setenv("BOT_USERNAME", "<your bot name>")
	os.Setenv("BOT_TOKEN", "<your bot token>")
	os.Setenv("CHANNEL_NAME", "<twitch channel name to join>")
}
```
