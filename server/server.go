package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func handleConnection(conn net.Conn) {
	request, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading client request:", err)
		return
	}

	// Extract username and password from the request
	parts := strings.Split(strings.TrimSpace(request), ",")
	username := parts[0]
	password := parts[1]

	// Check credentials
	if username == "std1" && password == "p@ssw0rd" {
		fmt.Fprint(conn, "Hello\n")
	} else {
		fmt.Fprint(conn, "Invalid credentials\n")
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting the server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on :8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		defer conn.Close()

		go handleConnection(conn)
	}
}
