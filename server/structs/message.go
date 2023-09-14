package structs

import (
	"chat-app/handlers"
	"strconv"
	"strings"
	"fmt"
)

type MessageCoreImpl struct {
	SendBy string
	Text   string
	Id     uint64
}

type Message struct {
	inner *MessageCoreImpl
}

func (m MessageCoreImpl) Marshal() string {
	return fmt.Sprintf(
		"%d %s:%s\n",
		m.Id,
		m.SendBy,
		m.Text,
	)
}

func (m *MessageCoreImpl) Unmarshal(s string) {
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

func (m MessageCoreImpl) Unique() string {
	return fmt.Sprint(m.Id)
}

//Wrapper for MessageCoreImpl

func (m *Message) Init(id uint64, send_by string, text string) {
	m.inner = &MessageCoreImpl {
		SendBy	: send_by,
		Text	: text,
		Id		: id,
	}
}

func (m Message) Marshal() string {
	return m.inner.Marshal()
}

func (m Message) Unmarshal(s string) {
	m.inner.Unmarshal(s)
}

func (m Message) Unique() string {
	return m.inner.Unique()
}

func (m Message) SendBy() any {
	return m.inner.SendBy
}

func (m *Message) SetText(text string) {
	m.inner.Text = text
}
