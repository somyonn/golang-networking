package main

import (
	"fmt"
	"net"
)

const (
	HOST = "127.0.0.1"
	PORT = "65456"
)

func main() {
	fmt.Println("> echo-server is activated")

	addr := net.UDPAddr{
		Port: 65456,
		IP:   net.ParseIP(HOST),
	}

	conn, _ := net.ListenUDP("udp", &addr)
	defer conn.Close()

	for {
		handleConnection(conn)
	}
}

func handleConnection(conn *net.UDPConn) {
	buffer := make([]byte, 1024)
	n, clientAddr, _ := conn.ReadFromUDP(buffer)
	recvData := buffer[:n]
	fmt.Printf("> echoed: %s\n", string(recvData))
	conn.WriteToUDP(recvData, clientAddr)
}
