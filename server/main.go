package main

import (
	"chat-app/connection"
	"chat-app/db"
	"chat-app/handlers"
	"chat-app/request"
	"chat-app/structs"
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

	conns := connection.BuildConnection(UserDb.Values(), addr)

	for {
		buf := make([]byte, 2048)
		_, addr, err := conn.ReadFrom(buf)
		s := string(buf)

		fmt.Printf("Request received from %s\n", addr.String())

		split := strings.Split(s, ";")
		split  = split[:len(split)-1]

		for _, q := range split {
			req := structs.Request {}
			req.Parse(q, addr)

			resp := request.Handle(req, conns)

			empty_user := structs.User {}
			empty_user.Init(nil, "")
	
			_, err = conn.WriteTo([]byte(resp), addr)
			handlers.HandleErr(err)
		}

		handlers.HandleErr(err)
	}
}
