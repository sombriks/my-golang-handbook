package main

import (
	"0015-rest-api/app"
	"log"
)

// application entrypoint
func main() {
	var s app.Server

	if err := s.Start(nil); err != nil {
		log.Fatal(err)
	}
}
