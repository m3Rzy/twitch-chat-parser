package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"twitch-chat-parser/config"
)

// Обработчик для корневого маршрута "/"
func homeHandler(w http.ResponseWriter, r *http.Request) {
	filePath := "../web/static/welcome.html"

	// Проверяем существование файла
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Println("Файл не найден:", filePath)
		http.Error(w, "Файл не найден", http.StatusNotFound)
		return
	}
	http.ServeFile(w, r, filePath)
}

// Структура для получения токена из JSON
type TokenRequest struct {
	AccessToken string `json:"access_token"`
}

// Обработчик для сохранения токена
func saveTokenHandler(w http.ResponseWriter, r *http.Request, tokenChannel chan<- string) {
	var req TokenRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Ошибка чтения данных", http.StatusBadRequest)
		return
	}

	// Логируем полученный токен и сохраняем его в конфиг
	config.TOKEN = req.AccessToken

	// Отправляем токен через канал
	tokenChannel <- req.AccessToken

	// Ответ клиенту
	fmt.Fprintf(w, "Токен успешно сохранен: %s", req.AccessToken)
}

// Основной обработчик
func MainHandler(tokenChannel chan<- string) {
	// Логируем текущую директорию
	_, err := os.Getwd()
	if err != nil {
		log.Fatal("Ошибка получения текущей директории:", err)
	}

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/save_token", func(w http.ResponseWriter, r *http.Request) {
		saveTokenHandler(w, r, tokenChannel) // Передаем канал в обработчик
	})

	log.Println("Сервер запущен на :8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Ошибка запуска сервера:", err)
	}
}
