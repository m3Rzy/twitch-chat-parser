package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Структура для получения токена и названия канала из JSON
type TokenRequest struct {
	AccessToken string `json:"access_token"`
	ChannelName string `json:"channel_name"`
}

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

// Обработчик для сохранения токена и названия канала
func saveTokenHandler(w http.ResponseWriter, r *http.Request, tokenChannel chan<- string, channel chan<- string) {
	var req TokenRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "Ошибка чтения данных", http.StatusBadRequest)
		return
	}
	// Отправляем токен через канал
	tokenChannel <- req.AccessToken
	channel <- req.ChannelName
	// Ответ клиенту
	fmt.Fprintf(w, "Токен и название канала успешно сохранены: %s, %s", req.AccessToken, req.ChannelName)
}

// Основной обработчик
func MainHandler(tokenChannel chan<- string, channel chan<- string) {
	// Логируем текущую директорию
	_, err := os.Getwd()
	if err != nil {
		log.Fatal("Ошибка получения текущей директории:", err)
	}

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/save_token", func(w http.ResponseWriter, r *http.Request) {
		saveTokenHandler(w, r, tokenChannel, channel) // Передаем канал в обработчик
	})
	http.HandleFunc("/success", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../web/static/success.html")
	})

	log.Println("Сервер запущен на :8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Ошибка запуска сервера:", err)
	}
}
