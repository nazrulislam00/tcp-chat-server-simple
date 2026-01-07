package main

import (
	"fmt"
	"net"
)

func main() {
	listener, _ := net.Listen("tcp", ":6000")
	fmt.Println("Server started on port 6000")

	conn, _ := listener.Accept()
	buffer := make([]byte, 1024)

	for {
		n, _ := conn.Read(buffer)
		message := string(buffer[:n])
		fmt.Println("Client:", message)

		if message == "exit" {
			break
		}

		conn.Write([]byte("Message received"))
	}

	conn.Close()
	listener.Close()
}
