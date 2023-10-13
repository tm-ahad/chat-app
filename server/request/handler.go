package request

import (
	"chat-app-server/db"
	"chat-app-server/enums/action"
	"chat-app-server/enums/response"
	"chat-app-server/interfaces"
	"chat-app-server/sa"
	"chat-app-server/structs"
	"chat-app-server/usernetwork"
)

var UserDb = db.NewDataBase("./storage/users.db")
var MsgDb  = db.NewDataBase("./storage/messages.db")

var idGen = structs.NewIdGen()

func MatchUserByAddr(usr interfaces.Model, match_to string) bool {
	u := usr.(structs.User)

	return u.Addr().String() == match_to
}

func Handle(req structs.Request, user_network usernetwork.UserNetwork) string {
	nil_msg := structs.Message {}
	nil_user := structs.User {}

	nil_msg.Init(0, "", "")
	nil_user.Init(nil, "")

	user_network.Insert(req.Conn, req.Addr)

	switch req.Action {
	case action.Say:
		obj 	:= UserDb.Find(req.Addr.String(), nil_user, MatchUserByAddr)
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

		MsgDb.Write(msg)
		user_network.SendToAllUsers([]byte(msg.Marshal()))

	case action.CreateUser:
		name := req.Param[0]
		obj  := UserDb.Find(name, nil_user, nil)

		if obj == nil {
			user := structs.User {}
			user.Init(req.Addr, name)

			UserDb.Write(user)
			return response.UserSaved
		} else {
			return response.UserExists
		}
		
		
	case action.RemoveUser:
		name := req.Param[0]
		rb   := UserDb.Find(name, nil_user, nil).(structs.User)

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
		obj := MsgDb.Find(id, nil_msg, nil)

		msg := structs.Message{}

		if obj == nil {
			return response.MessageNotFound
		} else {
			msg = obj.(structs.Message)
		}

		rb 	:= UserDb.Find(req.Addr.String(), nil_user, MatchUserByAddr).(structs.User)

		user_network.SendToAllUsers([]byte(req.RawBody))

		if msg.SendBy() == rb.Name() {
			MsgDb.Remove(id)
		} else {
			return response.AccessDenied
		}

	case action.ReplaceMsg:
		replace_with 	:= req.Param[1]
		id 				:= req.Param[0]

		obj 			:= MsgDb.Find(id, nil_msg, nil)
		msg				:= structs.Message {}

		if obj == nil {
			return response.MessageNotFound
		} else {
			msg = obj.(structs.Message)
		}

		rb 				:= UserDb.Find(req.Addr.String(), nil_user, MatchUserByAddr).(structs.User)

		user_network.SendToAllUsers([]byte(req.RawBody))

		if msg.SendBy() == rb.Name() {
			MsgDb.ReplaceMsgText(id, replace_with)
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
	
	return response.Empty
}
