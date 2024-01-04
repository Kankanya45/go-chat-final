package main

import (
	"fmt"
	"net"
	"strings"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	clientData := strings.TrimSpace(string(buffer[:n]))
	if isValidCredentials(clientData) {
		conn.Write([]byte("Hello"))
	} else {
		conn.Write([]byte("Invalid credentials"))
	}
}

func isValidCredentials(data string) bool {
	credentials := strings.Split(data, ":")
	if len(credentials) != 2 {
		return false
	}

	validUsername := "std1"
	validPassword := "p@ssw0rd"

	return credentials[0] == validUsername && credentials[1] == validPassword
}

func main() {
	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on :5000")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}
