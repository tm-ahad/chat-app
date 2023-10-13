package main

import (
	"chat-app-server/db"
	"chat-app-server/handlers"
	"chat-app-server/request"
	"chat-app-server/structs"
	"chat-app-server/usernetwork"
	"fmt"
	"log"
	"net"
	"strings"
)

var idGen = structs.NewIdGen() 
var addr  = ":9781"

var UserDb = db.NewDataBase("./storage/users.db")

func main() {
	log.SetFlags(log.Lshortfile)

	conn, err := net.ListenPacket("udp", ":9781")
	handlers.HandleErr(err)

	fmt.Println("Server started at http://127.0.0.1:9781")
	user_network := usernetwork.New(addr, UserDb.Values())

	for {
		buf := make([]byte, 2048)
		_, addr, err := conn.ReadFrom(buf)
		s := string(buf)

		fmt.Printf("Request received from %s\n", addr.String())

		split := strings.Split(s, ";")
		sl    := len(split)

		if sl > 1 {
			split = split[:sl-1]
		}

		for _, q := range split {
			req := structs.Request {}
			req.Parse(q, addr, conn)

			resp := request.Handle(req, user_network)

			empty_user := structs.User {}
			empty_user.Init(nil, "")
	
			_, err = conn.WriteTo([]byte(resp), addr)
			handlers.HandleErr(err)
		}

		handlers.HandleErr(err)
	}
}
