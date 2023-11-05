package main

import (
	"fmt"
	"net"
)

func EchoUDPServer(address string) {
	// not handling errors for simplicity
	addr, _ := net.ResolveUDPAddr("udp", address)
	connection, _ := net.ListenUDP("udp", addr)
	defer connection.Close()
	// Two key differences for UDP:
	// 1- we must know the payload size
	// 2- we need to get sender address, so we can answer
	var line [64]byte
	_, sender, _ := connection.ReadFromUDP(line[:])
	echo := fmt.Sprintf("You said: %s\n", string(line[:]))
	connection.WriteToUDP([]byte(echo), sender)
}
