package request

import (
	"strings"
	"net"
)

type Request struct {
	Param  []string
	Addr   net.Addr
	Action string
}

func (req *Request) Parse(s string, addr net.Addr)  {
	split := strings.Split(s, ":")
	action := split[0]

	req.Action = action
	req.Addr = addr

	rem := split[1:]

	for _, m := range rem {
		req.Param = append(req.Param, m)
	}
}
