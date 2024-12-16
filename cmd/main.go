package main

import (
	"twitch-chat-parser/config"
	"twitch-chat-parser/internal/services"
)

func main() {
	conn := services.IrcConnection(config.GetConfigs())
	defer conn.Close()

	for {
		message := services.ParseChatToString(conn)
		services.HandleChatMessage(message)
	}
}
