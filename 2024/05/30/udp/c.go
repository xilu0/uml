package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "localhost:8080")
	if err != nil {
		fmt.Println("Error resolving address:", err)
		return
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("Error dialing:", err)
		return
	}
	defer conn.Close()

	message := []byte("Hello, UDP Server!")
	_, err = conn.Write(message)
	if err != nil {
		fmt.Println("Error writing to UDP server:", err)
		os.Exit(1)
	}

	buffer := make([]byte, 1024)
	n, _, err := conn.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println("Error reading response from UDP server:", err)
		return
	}

	fmt.Printf("Received response from server: %s\n", string(buffer[:n]))
}
