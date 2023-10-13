package structs

import (
	"chat-app-server/handlers"
	"fmt"
	"net"
	"strings"
)

type UserCoreImpl struct {
	Addr net.Addr
	Name string
}

type User struct {
	inner *UserCoreImpl
}

func (user UserCoreImpl) Marshal() string {
	return fmt.Sprintf(
		"%s:%s",
		user.Name,
		user.Addr.String(),
	)
}

func (user *UserCoreImpl) Unmarshal(s string) {
	split 		:= strings.Split(s, ":")
	raw_addr 	:= strings.Join(split[1:], ":")

	addr, err := net.ResolveUDPAddr("udp", raw_addr)

	handlers.HandleErr(err)

	user.Name = split[0]
	user.Addr = addr
}

func (user UserCoreImpl) Unique() string {
	return user.Name
}

//Wrapper for UserCoreImpl

func (user *User) Init(addr net.Addr, s string) {
	user.inner = &UserCoreImpl {
		Addr: addr,
		Name: s,
	}
}

func (user User) Marshal() string {
	return user.inner.Marshal()
}

func (user User) Unmarshal(s string) {
	user.inner.Unmarshal(s)
}

func (user User) Unique() string {
	return user.inner.Unique()
}

func (user User) Name() string {
	return user.inner.Name
}

func (user User) Addr() net.Addr {
	return user.inner.Addr
}
