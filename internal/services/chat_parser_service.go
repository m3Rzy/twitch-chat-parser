package services

import (
    "bufio"
    "fmt"
    "net"
    "strings"
)

func ParseChatToString(conn net.Conn) string {
    reader := bufio.NewReader(conn)

    for {
        message, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("Ошибка чтения сообщения:", err)
            return ""
        }

        // Убираем пробелы и переносы строк
        message = strings.TrimSpace(message)

        // Обработка PING от сервера (обязательно для поддержания соединения)
        if strings.HasPrefix(message, "PING") {
            fmt.Fprintf(conn, "PONG :tmi.twitch.tv\r\n")
            continue
        }

        return message
    }
}
