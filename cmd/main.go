package main

import (
	"log"
	"twitch-chat-parser/config"
	"twitch-chat-parser/internal/services"
	"twitch-chat-parser/internal/transport/rest"
	"time"
)

func main() {
	// Канал для уведомления о получении токена
	tokenChannel := make(chan string)

	// Запуск HTTP-сервера в отдельной горутине
	go func() {
		log.Println("Запускаем HTTP-сервер...")
		rest.MainHandler(tokenChannel) // Передаем канал в обработчик
	}()

	// Ожидаем, пока не получим токен через канал
	token := <-tokenChannel
	log.Printf("Получен токен: %s", token)

	conn := services.IrcConnection(config.GetConfigs())
	defer conn.Close()

	for {
		message := services.ParseChatToString(conn)
		services.HandleChatMessage(message)

		time.Sleep(1 * time.Second)
	}
}
