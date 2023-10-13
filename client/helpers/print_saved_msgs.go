package helpers

import (
	"chat-app-client/handlers"
	"fmt"
	"net"
)

func Print_saved_msgs(conn net.Conn) {
    _, err := conn.Write([]byte("get_msgs;"))
    handlers.HandleErr(err)

    received := make([]byte, 1024)
    _, er := conn.Read(received)
    handlers.HandleErr(er)

    fmt.Println(string(received))
}
