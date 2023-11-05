package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func EchoUDPClient(address string, message string) string {
	// not handling errors for simplicity
	connection, _ := net.Dial("udp", address)
	defer connection.Close()
	io.WriteString(connection, fmt.Sprintf("%s\n", message))
	lineReader := bufio.NewReader(connection)
	line, _ := lineReader.ReadString('\n')
	return line
}
