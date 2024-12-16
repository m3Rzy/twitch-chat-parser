package services

import (
	"fmt"
	"strings"
)

func HandleChatMessage(message string) {

	if !strings.Contains(message, "PRIVMSG") {
		return
	}

	// Пример сообщения:
	// :username!username@username.tmi.twitch.tv PRIVMSG #channel :Hello, chat!
	parts := strings.Split(message, "PRIVMSG")
	if len(parts) < 2 {
		return
	}

	// Извлекаем пользователя и текст сообщения
	user := strings.Split(parts[0], "!")
	if len(user) < 2 {
		return
	}
	username := user[0][1:] // Убираем первый символ (:)

	// Извлекаем текст сообщения
	msg := strings.SplitN(parts[1], ":", 2)
	if len(msg) < 2 {
		return
	}
	text := msg[1]

	// Выводим в консоль
	fmt.Printf("[%s]: %s\n", username, text)
}