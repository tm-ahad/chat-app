package main

import (
	"chat-app/connection"
	"chat-app/db"
	"chat-app/enums"
	"chat-app/handlers"
	"chat-app/structs"
	"fmt"
	"log"
	"net"
	"strings"
)

var idGen = structs.NewIdGen()
var addr  = ":9781"

var UserDb = db.NewDataBase("./storage/users.db")
var MsgDb = db.NewDataBase("./storage/messages.db")

func handleRequest(req structs.Request, conns []net.Conn) string {
	switch req.Action {
	case action.Say:
		empty_user := structs.User{}
		empty_user.Init(nil, "")

		obj 	:= UserDb.Find(req.Addr.String(), empty_user);
		user 	:= structs.User{}

		if obj == nil {
			return "User not found"
		} else {
			user = obj.(structs.User)
		}

		text := req.Param[0]
		var msg = structs.Message {
			SendBy: user.Name(),
			Text:   text,
			Id:     idGen.Gen(),
		}
		string_msg  := fmt.Sprintf("%s:%s", user.Name(), msg.Text)

		MsgDb.Write(msg)
		connection.Send(conns, string_msg)

		return "Message saved successfully!"
	case action.SetUser:
		name := req.Param[0]
		user := structs.User{}

		user.Init(req.Addr, name)

		UserDb.Write(user)
		return "User saved."
	}

	return "Invalid action."
}

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
			req := structs.Request{}
			req.Parse(q, addr)

			resp := handleRequest(req, conns)

			empty_user := structs.User{}
			empty_user.Init(nil, "")
	
			_, err = conn.WriteTo([]byte(resp), addr)
			handlers.HandleErr(err)
		}

		handlers.HandleErr(err)
	}
}
