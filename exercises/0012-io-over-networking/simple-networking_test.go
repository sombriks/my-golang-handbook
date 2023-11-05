package main

import (
	"testing"
	"time"
)

func TestEchoTCP(t *testing.T) {
	addr := ":8765"
	message := "Hail from the other side!"

	go EchoTCPServer(addr)

	// must respect the time needed by the OS to provision the resource
	time.Sleep(time.Duration(1) * time.Second)

	result := EchoTCPClient(addr, message)

	if "You said: Hail from the other side!\n" != result {
		t.Fatal("Echoed message isn't the expected one")
	}

}

func TestEchoUDP(t *testing.T) {
	addr := ":9876"
	message := "Hail from the other side!"

	go EchoUDPServer(addr)

	time.Sleep(time.Duration(1) * time.Second)

	result := EchoUDPClient(addr, message)

	if "You said: Hail from the other side!\n" != result {
		t.Fatal("Echoed message isn't the expected one")
	}
}
