package handlers

import "log"

func HandleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
