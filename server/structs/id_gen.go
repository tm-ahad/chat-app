package structs

import (
	"chat-app/handlers"
	"strconv"
	"strings"
	"syscall"
	"errors"
	"fmt"
	"io"
	"os"
)

type IdGen struct {
	file *os.File
}

func NewIdGen() IdGen {
	file, err := os.OpenFile("./storage/id", syscall.O_RDWR, 0666)
	handlers.HandleErr(err)

	return IdGen {
		file: file,
	}
}

func (inst *IdGen) Gen() uint64 {
	s, _ := inst.file.Stat()
	f_len := s.Size()
	f_len++

	b := make([]byte, f_len)

	for {
		_, err := inst.file.Read(b)

		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			handlers.HandleErr(err)
		}
	}

	var id strings.Builder

	for _, byte_v := range b {
		if byte_v > 0 {
			id.WriteByte(byte_v)
		}
	}

	numb, err := strconv.ParseUint(id.String(), 10, 64)
	handlers.HandleErr(err)

	handlers.HandleErr(inst.file.Truncate(0))
	_, e := inst.file.Seek(0, 0)

	handlers.HandleErr(e)

	uid := []byte(fmt.Sprint(numb + 1))
	_, er := inst.file.WriteAt(uid, 0)

	handlers.HandleErr(er)

	return numb
}
