package main

import (
	"log"
	"twitch-chat-parser/internal/services"
	"twitch-chat-parser/internal/transport/rest"
)

var USERNAME = "default"

func main() {
	// Канал для уведомления о получении токена
	tokenChannel := make(chan string)
	channel := make(chan string)

	go func() {
		log.Println("Запускаем HTTP-сервер...")
		rest.MainHandler(tokenChannel, channel)
	}()

	token := <-tokenChannel
	log.Printf("Получен токен: %s", token)
	channel_name := <-channel
	log.Printf("Найден канал стримера: %s", channel_name)

	conn := services.IrcConnection("oauth:" + token, "#" + channel_name)
	defer conn.Close()

	for {
		message := services.ParseChatToString(conn)
		services.HandleChatMessage(message)
	}
}
