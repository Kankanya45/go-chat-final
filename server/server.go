package main

import (
	"fmt"
	"net"
	"strings"
)

const (
	port          = 5000
	validUsername = "std1"
	validPassword = "p@ssw0rd"
	bufferSize    = 1024
	helloResponse = "Hello"
	errorResponse = "Invalid credentials"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, bufferSize)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	clientData := strings.TrimSpace(string(buffer[:n]))
	if isValidCredentials(clientData) {
		conn.Write([]byte(helloResponse))
	} else {
		conn.Write([]byte(errorResponse))
	}
}

func isValidCredentials(data string) bool {
	credentials := strings.Split(data, ":")
	if len(credentials) != 2 {
		return false
	}

	return credentials[0] == validUsername && credentials[1] == validPassword
}

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	fmt.Printf("Server is listening on port %d\n", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}
