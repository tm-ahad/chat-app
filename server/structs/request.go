package structs

import (
	"chat-app-server/enums/action"
	"net"
	"strings"
)

func RemoveNullBytes(s string) string {
	bytes  := []byte(s)
	sb     := strings.Builder {}

	for _, b := range bytes {
		if b != 0 {
			sb.WriteByte(b)
		} else {
			break
		}
	}

	return sb.String()
}


type Request struct {
	Conn    net.PacketConn
	Param   []string
	Addr    net.Addr
	Action  string
	RawBody string
}

func (req *Request) Parse(s string, addr net.Addr, conn net.PacketConn)  {
	split := strings.Split(s, ":")
	s0    := split[0]

	req.Addr    = addr
	req.RawBody = s
	req.Conn    = conn

	if s0 == s && s0 != "get_msgs" && s0 != "get_users" {
		s = RemoveNullBytes(s)
		req.Action = action.Say
		req.Param  = append(req.Param, s)
	} else {
		action    := split[0]
		req.Action = action

		rem := split[1:]

		for _, m := range rem {
			m = RemoveNullBytes(m)
			req.Param = append(req.Param, m)
		}
	}
}
