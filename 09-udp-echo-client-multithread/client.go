package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	HOST = "127.0.0.1"
	PORT = "65456"
)

func main() {
	fmt.Println("> echo-client is activated")

	serverAddr, _ := net.ResolveUDPAddr("udp", HOST+":"+PORT)
	conn, _ := net.DialUDP("udp", nil, serverAddr)
	defer conn.Close()

	go recvHandler(conn)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		sendMsg, _ := reader.ReadString('\n')
		sendMsg = strings.TrimSpace(sendMsg)
		conn.Write([]byte(sendMsg))
		if sendMsg == "quit" {
			break
		}
	}

	fmt.Println("> echo-client is de-activated")
}

func recvHandler(conn *net.UDPConn) {
	buffer := make([]byte, 1024)
	for {
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			break
		}
		recvData := string(buffer[:n])
		fmt.Println("> received:", recvData)
		if recvData == "quit" {
			break
		}
	}
}
