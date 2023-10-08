package db

import (
	"chat-app/handlers"
	"chat-app/helpers"
	"chat-app/interfaces"
	"chat-app/structs"
	"fmt"
	"os"
	"strings"
	"syscall"
)

type DataBase struct {
	cache  		map[any]interfaces.Model
	keys 		[]string
	values 		[]string
	file 		*os.File
	cont 		string
}

func (db DataBase) Write(obj interfaces.Model) {
	_, err := db.file.WriteString(obj.Marshal())

	handlers.HandleErr(err)
}

func (db DataBase) RawCont() string {
	return db.cont
}

func (db DataBase) Find(key string, obj interfaces.Model) interfaces.Model {
	val := db.cache[key]

	if val == nil {
		split := strings.Split(db.cont, "\n")

		for _, line := range split {
			if len(line) == 0 {
				continue
			}

			obj.Unmarshal(line)

			if obj.Unique() == key {
				db.cache[key] = obj
				return obj
			}	
		}
	} else {
		return val
	}

	return nil
}

func (db DataBase) Remove(key any) {
	split 			:= strings.Split(db.cont, "\n")
	updated_cont 	:= strings.Builder{}

	for _, line := range split {
		line = fmt.Sprintf("%s\n", line)
		if !strings.HasPrefix(line, fmt.Sprint(key)) {
			updated_cont.WriteString(line)
		}
	}

	s := updated_cont.String()
	handlers.HandleErr(db.file.Truncate(0))

	_, err := db.file.Seek(0, 0)
	handlers.HandleErr(err)
	
	_, er := db.file.WriteString(s)
	handlers.HandleErr(er)

	db.cache[key] = nil
	db.cont = s
}

func (db DataBase) ReplaceMsgText(key any, update_with string) {
	split 			:= strings.Split(db.cont, "\n")
	updated_cont 	:= strings.Builder {}

	for _, line := range split {
		line = fmt.Sprintf("%s\n", line)

		if strings.HasPrefix(line, key.(string)) {
			col 		:= uint(strings.Index(line, ":"))
			rng 		:= structs.NewRange(col+1, uint(len(line)))

			upd_line 	:= helpers.ReplaceRange(line, rng, update_with)

			updated_cont.WriteString(upd_line)
		} else {
			updated_cont.WriteString(line)
		}
	}

	s := updated_cont.String()
	handlers.HandleErr(db.file.Truncate(0))

	_, err := db.file.Seek(0, 0)
	handlers.HandleErr(err)

	_, er := db.file.WriteString(s)
	handlers.HandleErr(er)

	val := db.cache[key]
	msg := val.(structs.Message)

	if val != nil {
		db.cache[key] = nil
		msg.SetText(update_with)

		db.cache[key] = msg
	}
	
	db.cont = s
}

func (db DataBase) Values() []string {
	return db.values
}

func (db DataBase) Keys() []string {
	return db.keys
}

func NewDataBase(path string) DataBase {
	file, err := os.OpenFile(path, syscall.O_RDWR, 0666)
	handlers.HandleErr(err)

	s, _ := file.Stat()
	f_len := s.Size()

	b := make([]byte, f_len)

	_, er := file.Read(b)
	handlers.HandleErr(er)

	cont         := string(b)
	keys, values := helpers.KeysAndValues(cont)

	return DataBase {
		values: 	values,
		keys:		keys,
		cache:  	make(map[any]interfaces.Model),
		file: 		file,
		cont: 		cont,
	}
}
