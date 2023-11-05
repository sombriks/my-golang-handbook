package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func EchoClient(address string, message string) string {
	connection, errConnection := net.Dial("tcp", address)
	if errConnection != nil {
		log.Panic(errConnection)
		return ""
	}
	defer connection.Close()

	_, errResponse := io.WriteString(connection, fmt.Sprintf("%s\n", message))
	if errResponse != nil {
		log.Panic(errResponse)
		return ""
	}

	lineReader := bufio.NewReader(connection)
	line, errLine := lineReader.ReadString('\n')
	if errLine != nil {
		log.Panic(errLine)
		return ""
	}
	return line
}
