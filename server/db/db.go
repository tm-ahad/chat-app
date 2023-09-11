package db

import (
	"chat-app/handlers"
	"chat-app/helpers"
	"chat-app/interfaces"
	"os"
	"strings"
	"syscall"
)

type DataBase struct {
	file 	*os.File
	cont 	string
	values 	[]string
}

func (db DataBase) Write(obj interfaces.Model) {
	_, err := db.file.WriteString(obj.Marshal())

	handlers.HandleErr(err)
}

func (db DataBase) Find(key any, obj interfaces.Model) interfaces.Model {
	split := strings.Split(db.cont, "\n")

	for _, line := range split {
		if len(line) == 0 {
			continue
		}

		obj.Unmarshal(line)

		if obj.Unique() == key {
			return obj
		}
	}

	return nil
}

func (db DataBase) Values() []string {
	return db.values
}

func NewDataBase(path string) DataBase {
	file, err := os.OpenFile(path, syscall.O_RDWR, 0666)
	handlers.HandleErr(err)

	s, _ := file.Stat()
	f_len := s.Size()

	b := make([]byte, f_len)

	_, er := file.Read(b)
	handlers.HandleErr(er)

	cont := string(b)

	return DataBase {
		file: 	file,
		cont: 	cont,
		values: helpers.Values(cont),
	}
}
