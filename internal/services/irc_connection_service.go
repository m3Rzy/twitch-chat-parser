package services

import (
	"fmt"
	"log"
	"net"
)

var USERNAME = "default"

func IrcConnection(oauthToken, channel string) net.Conn {
	conn, err := net.Dial("tcp", "irc.chat.twitch.tv:6667")
	if err != nil {
		log.Fatal("Error connecting to IRC server: ", err)
	}

	fmt.Fprintf(conn, "PASS %s\r\n", oauthToken)
	fmt.Fprintf(conn, "NICK %s\r\n", USERNAME)
	fmt.Fprintf(conn, "JOIN %s\r\n", channel)

	return conn
}
