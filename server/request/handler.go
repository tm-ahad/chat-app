package request

import (
	"chat-app/connection"
	"chat-app/db"
	"chat-app/enums/action"
	"chat-app/enums/response"
	"chat-app/structs"
	"net"
)

var UserDb = db.NewDataBase("./storage/users.db")
var MsgDb = db.NewDataBase("./storage/messages.db")

var idGen = structs.NewIdGen()

func Handle(req structs.Request, conns []net.Conn) string {
	switch req.Action {
	case action.Say:
		empty_user := structs.User{}
		empty_user.Init(nil, "")

		obj 	:= UserDb.Find(req.Addr.String(), empty_user);
		user 	:= structs.User{}

		if obj == nil {
			return response.UserNotFound  
		} else {
			user = obj.(structs.User)
		}

		user_name := user.Name()

		text := req.Param[0]
		var msg = structs.Message {
			SendBy: user_name,
			Text:   text,
			Id:     idGen.Gen(),
		}
		string_msg  := response.Marshal(action.Say, []string{user_name, text})

		MsgDb.Write(msg)
		connection.Send(conns, string_msg)

		return response.MessageSaved
	case action.CreateUser:
		name := req.Param[0]
		user := structs.User{}

		user.Init(req.Addr, name)

		UserDb.Write(user)
		return "User saved."
	case action.RemoveUser:
		name := req.Param[0]

		empty_user := structs.User{}
		empty_user.Init(nil, "");

		rb := UserDb.Find(req.Addr.String(), empty_user).(structs.User)

		if rb.Name() == name {
			user := structs.User{}
			user.Init(req.Addr, name)

			UserDb.Remove(name)
			return response.UserRemoved
		} else {
			return "Access denied."
		}

		
	case action.RemoveMsg:
		id 	:= req.Param[0]
		msg := MsgDb.Find(id, structs.Message{}).(structs.Message)
		rb 	:= UserDb.Find(req.Addr.String(), structs.User{}).(structs.User)

		if rb.Name() == msg.SendBy {
			MsgDb.Remove(id)
			return response.MessageRemoved
		} else {
			return response.AccessDenied
		}

	case action.ReplaceMsg:
		id 				:= req.Param[0]
		replace_with 	:= req.Param[1]

		msg 			:= MsgDb.Find(id, structs.Message{}).(structs.Message)
		rb 				:= UserDb.Find(req.Addr.String(), structs.User{}).(structs.User)

		if rb.Name() == msg.SendBy {
			MsgDb.ReplaceMsgText(id, replace_with)
			return response.MessageRemoved
		} else {
			return response.AccessDenied
		}
	}

	return "Invalid action."
}