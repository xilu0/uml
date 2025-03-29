package main

import (
	"fmt"
	"net"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", ":8080")
	if err != nil {
		fmt.Println("Error resolving address:", err)
		return
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer conn.Close()

	buffer := make([]byte, 1024)
	for {
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading from UDP client:", err)
			continue
		}

		fmt.Printf("Received message from %s: %s\n", clientAddr, string(buffer[:n]))

		response := []byte("Message received")
		_, err = conn.WriteToUDP(response, clientAddr)
		if err != nil {
			fmt.Println("Error sending response to UDP client:", err)
		}
	}
}
