package main

import (
	"os"
	"fmt"
	"net"
	"strings"
)

const END_DATA = "[END-DATA]"

func main() {
	listen, err := net.Listen("tcp", ":8080")
	defer listen.Close()
	if err != nil {
		os.Exit(1)
	}
	fmt.Println("Server is listening...")
	for {
		conn, err := listen.Accept()
		if err != nil { continue }
		go server(conn)
	}
}

func server(conn net.Conn) {
	defer conn.Close()

	var buffer = make([]byte, 256)
	var message string

	for {
		lenght, err := conn.Read(buffer)
		if err != nil {
			return
		}
		message += string(buffer[:lenght])
		if strings.HasSuffix(message, END_DATA) {
			message = strings.TrimSuffix(message, END_DATA)
			break
		}
	}
	conn.Write([]byte(strings.ToUpper(message)))
}