package main

import (
	"0015-rest-api/app"
	"log"
)

func main() {
	var s app.Server
	err := s.Start()
	if err != nil {
		log.Fatal(err)
	}
}
