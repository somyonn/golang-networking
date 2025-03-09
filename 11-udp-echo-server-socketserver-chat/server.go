package main

import (
	"fmt"
	"net"
	"strings"
	"sync"
)

var clients = make([]*net.UDPAddr, 0)
var mutex sync.Mutex

const (
	HOST = "127.0.0.1"
	PORT = "65456"
)

func main() {
	fmt.Println("> echo-server is activated")
	addr := net.UDPAddr{Port: 65456, IP: net.ParseIP(HOST)}
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Println("> listen() failed and program terminated:", err)
		return
	}
	defer conn.Close()

	for {
		handleConnection(conn)
	}
}

func handleConnection(conn *net.UDPConn) {
	buffer := make([]byte, 1024)
	n, clientAddr, err := conn.ReadFromUDP(buffer)
	if err != nil {
		return
	}

	recvData := strings.TrimSpace(string(buffer[:n]))

	if strings.HasPrefix(recvData, "#") || recvData == "quit" {
		if recvData == "#REG" {
			fmt.Println("> client registered", clientAddr)
			mutex.Lock()
			clients = append(clients, clientAddr)
			mutex.Unlock()
		} else if recvData == "#DEREG" || recvData == "quit" {
			mutex.Lock()
			for i, addr := range clients {
				if addr.String() == clientAddr.String() {
					fmt.Println("> client de-registered", clientAddr)
					clients = append(clients[:i], clients[i+1:]...)
					break
				}
			}
			mutex.Unlock()
		}
	} else {
		mutex.Lock()
		if len(clients) == 0 {
			fmt.Println("> no clients to echo")
		} else {
			registered := false
			for _, addr := range clients {
				if addr.String() == clientAddr.String() {
					registered = true
					break
				}
			}
			if !registered {
				fmt.Println("> ignores a message from un-registered client")
			} else {
				fmt.Printf("> received (%s) and echoed to %d clients\n", recvData, len(clients))
				for _, addr := range clients {
					conn.WriteToUDP([]byte(recvData), addr)
				}
			}
		}
		mutex.Unlock()
	}
}
