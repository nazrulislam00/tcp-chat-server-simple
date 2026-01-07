package main

import (
	"fmt"
	"net"
)

func main() {
	conn, _ := net.Dial("tcp", "localhost:6000")
	buffer := make([]byte, 1024)

	for {
		fmt.Print("You: ")
		var msg string
		fmt.Scanln(&msg)

		conn.Write([]byte(msg))
		n, _ := conn.Read(buffer)
		fmt.Println("Server:", string(buffer[:n]))

		if msg == "exit" {
			break
		}
	}

	conn.Close()
}
