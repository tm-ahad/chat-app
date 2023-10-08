package request

import (
	"chat-app/connection"
	"chat-app/db"
	"chat-app/enums/action"
	"chat-app/enums/response"
	"chat-app/sa"
	"chat-app/structs"
	"fmt"
	"net"
)

var UserDb = db.NewDataBase("./storage/users.db")
var MsgDb  = db.NewDataBase("./storage/messages.db")

var idGen = structs.NewIdGen()

func Handle(req structs.Request, conns []net.Conn) string {
	nil_msg := structs.Message {}
	nil_user := structs.User {}

	nil_msg.Init(0, "", "")
	nil_user.Init(nil, "")

	fmt.Printf("%+v\n", req)

	switch req.Action {
	case action.Say:
		obj 	:= UserDb.Find(req.Addr.String(), nil_user);
		user 	:= structs.User{}

		if obj == nil {
			return response.UserNotFound  
		} else {
			user = obj.(structs.User)
		}

		user_name := user.Name()
		text      := req.Param[0]

		var msg = structs.Message {}
		msg.Init(idGen.Gen(), user_name, text)

		string_msg  := response.Marshal(action.Say, []string{user_name, text})

		MsgDb.Write(msg)
		connection.Send(conns, string_msg)

		return response.MessageSaved

	case action.CreateUser:
		name := req.Param[0]
		user := structs.User {}

		user.Init(req.Addr, name)

		UserDb.Write(user)
		return response.UserSaved
		
	case action.RemoveUser:
		name := req.Param[0]
		rb := UserDb.Find(req.Addr.String(), nil_user).(structs.User)

		if rb.Name() == name {
			user := structs.User {}
			user.Init(req.Addr, name)

			UserDb.Remove(name)
			return response.UserRemoved
		} else {
			return response.AccessDenied
		}

		
	case action.RemoveMsg:
		id := req.Param[0]

		obj := MsgDb.Find(id, nil_msg);
		msg := structs.Message{}

		if obj == nil {
			return response.MessageNotFound
		} else {
			msg = obj.(structs.Message)
		}

		rb 	:= UserDb.Find(req.Addr.String(), nil_user).(structs.User)
		connection.Send(conns, req.RawBody)

		if msg.SendBy() == rb.Name() {
			MsgDb.Remove(id)
			return response.MessageRemoved
		} else {
			return response.AccessDenied
		}

	case action.ReplaceMsg:
		replace_with 	:= req.Param[1]
		id 				:= req.Param[0]

		obj 			:= MsgDb.Find(id, nil_msg)
		msg				:= structs.Message {}

		if obj == nil {
			return response.MessageNotFound
		} else {
			msg = obj.(structs.Message)
		}

		rb 				:= UserDb.Find(req.Addr.String(), nil_user).(structs.User)

		connection.Send(conns, req.RawBody)

		if msg.SendBy() == rb.Name() {
			MsgDb.ReplaceMsgText(id, replace_with)
			return response.MessageRemoved
		} else {
			return response.AccessDenied
		}

	case action.GetUsers:
		user_names := UserDb.Keys()
		return sa.Marshal(user_names)

	case action.GetMsgs:
		msgs := MsgDb.RawCont()
		return msgs

	default:
		return "Invalid action."
	}
}
