package structs

import (
	"chat-app/handlers"
	"strconv"
	"strings"
	"fmt"
)

type Message struct {
	SendBy string
	Id     uint64
	Text   string
}

func (m Message) Marshal() string {
	return fmt.Sprintf(
		"%d %s:%s\n",
		m.Id,
		m.SendBy,
		m.Text,
	)
}

func (m Message) Unmarshal(s string) {
	split 		:= strings.Split(s, " ")
	id, e 	    := strconv.ParseUint(split[0], 10, 64)
	handlers.HandleErr(e)

	rem 		:= split[1]

	split2 		:= strings.Split(rem, ":")
	user_name 	:= split2[0]
	msg_text  	:= split2[1]

	m.Id = id;
	m.SendBy = user_name;
	m.Text = msg_text
}

func (m Message) Unique() any {
	return m.Id
}
