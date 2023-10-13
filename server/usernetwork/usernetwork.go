package usernetwork

import (
	"chat-app-server/handlers"
	"net"
)

type UserNetwork struct {
	conns map[net.Addr]net.PacketConn
	laddr *net.UDPAddr
}

func New(laddr string, user_ips []string) UserNetwork {
	_laddr, err := net.ResolveUDPAddr("udp", laddr)
	handlers.HandleErr(err)

	return UserNetwork {
		laddr: _laddr,
		conns: make(map[net.Addr]net.PacketConn),
	}
}

func (un *UserNetwork) SendToAllUsers(msg []byte)  {
	for addr, conn := range un.conns {
		conn.WriteTo(msg, addr)
	}
}

func (un *UserNetwork) Insert(conn net.PacketConn, addr net.Addr) {
	un.conns[addr] = conn
}
