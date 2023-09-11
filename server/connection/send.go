package connection

import "net"

func Send(conns []net.Conn, msg string) {
	for _, conn := range conns {
		conn.Write([]byte(msg))
	}
}
