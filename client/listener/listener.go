package listener

import (
	"fmt"
	"net"
)

func Listen(server net.Conn) {
	for {
		buf := make([]byte, 1024)
		_, err := server.Read(buf)
		body := string(buf)

		fmt.Println(body)

		if err != nil {
			continue
		}
	}
}
