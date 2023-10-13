package main

import (
	"chat-app-client/handlers"
	"chat-app-client/helpers"
	"chat-app-client/listener"
	"fmt"
	"net"
)

func main() {
	var account_number int

	fmt.Println("This number is a password. This number is needed to log in to your account")
	fmt.Printf("Your account number: ")
	fmt.Scanf("%d", &account_number)

	chatServer, err := net.ResolveUDPAddr("udp", ":9781")
	handlers.HandleErr(err)

	server_addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", account_number))
	handlers.HandleErr(err)

	conn, err := net.DialUDP("udp", server_addr, chatServer)
	handlers.HandleErr(err)
	helpers.Print_saved_msgs(conn)

	go listener.Listen(conn)

	for {
		var input string
		fmt.Scanf("%s", &input)

		_, err = conn.Write([]byte(input))
		handlers.HandleErr(err)

		received := make([]byte, 1024)
		_, err = conn.Read(received)
		handlers.HandleErr(err)

		if len(received) != 0 {
			fmt.Printf("%s\n", string(received))
		}
	}
}
