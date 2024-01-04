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
		fmt.Println("Error reading data:", err)
		return
	}

	clientData := strings.TrimSpace(string(buffer[:n]))
	if isValidCredentials(clientData) {
		conn.Write([]byte("Hello"))
	} else {
		conn.Write([]byte("Invalid credentials"))
	}
}

// ตรวจสอบข้อมูลบัญชีผู้ใช้
func isValidCredentials(data string) bool {
	// กำหนดข้อมูลบัญชีผู้ใช้ที่ถูกต้อง
	validCredentials := "std1:p@ssw0rd"

	return data == validCredentials
}

func main() {
	listener, err := net.Listen("tcp", ":5000") //เชื่อมต่อ TCP บนพอร์ต 5000 และรอรับการเชื่อมต่อ:
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
