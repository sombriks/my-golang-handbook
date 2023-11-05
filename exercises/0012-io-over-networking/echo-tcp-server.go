package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func EchoTCPServer(address string) {
	listener, errServer := net.Listen("tcp", address)
	if errServer != nil {
		log.Panic(errServer)
		return
	}
	defer listener.Close()

	//for {
	connection, errConnection := listener.Accept()
	if errConnection != nil {
		log.Panic(errConnection)
		return
	}
	lineReader := bufio.NewReader(connection)
	line, errLine := lineReader.ReadString('\n')
	if errLine != nil {
		log.Panic(errLine)
		return
	}
	_, errResponse := io.WriteString(connection, fmt.Sprintf("You said: %s\n", line))
	if errResponse != nil {
		log.Panic(errResponse)
		return
	}
	connection.Close()
	//}
}
