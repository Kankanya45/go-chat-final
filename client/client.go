package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Print("Enter username: ")
	username, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Enter password: ")
	password, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	password = strings.TrimSpace(password)

	connection, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to the server:", err)
		return
	}
	defer connection.Close()

	message := fmt.Sprintf("%s,%s", username, password)
	fmt.Fprintf(connection, message+"\n")

	response, err := bufio.NewReader(connection).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading server response:", err)
		return
	}

	fmt.Println("Server response:", response)
}
