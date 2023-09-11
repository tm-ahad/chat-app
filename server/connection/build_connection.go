package connection

import (
	"chat-app/handlers"
	"net"
)
 
func BuildConnection(addrs []string, curr_addr string) []net.Conn {
	laddr, err := net.ResolveUDPAddr("udp", curr_addr)
	handlers.HandleErr(err)

	var conns []net.Conn

	for _, raw_addr := range addrs {
		addr, err := net.ResolveUDPAddr("udp", raw_addr);
		if err != nil {continue}

		conn, er := net.DialUDP("udp", laddr, addr)
		handlers.HandleErr(er)

		conns = append(conns, conn)
	}

	return conns
}
