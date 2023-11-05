package main

import (
	"testing"
	"time"
)

func TestSimpleEcho(t *testing.T) {
	addr := ":8765"
	message := "Hail from the other side!"

	go EchoServer(addr)

	// must respect the time needed by the OS to provision the resource
	time.Sleep(time.Duration(1) * time.Second)

	result := EchoClient(addr, message)

	if "You said: Hail from the other side!\n" != result {
		t.Fatal("Echoed message isn't the expected one")
	}

}
