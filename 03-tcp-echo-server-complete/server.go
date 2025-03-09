package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	fmt.Println("> echo-server is activated")
	listener, err := net.Listen("tcp", "127.0.0.1:65456")
	if err != nil {
		fmt.Println("> bind() failed and program terminated")
		return
	}
	defer listener.Close()

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("> listen() failed and program terminated")
		return
	}
	defer conn.Close()

	clientAddress := conn.RemoteAddr().(*net.TCPAddr)
	fmt.Printf("> client connected by IP address %s with Port number %d\n", clientAddress.IP.String(), clientAddress.Port)
	reader := bufio.NewReader(conn)
	for {
		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		fmt.Println("> echoed:", msg)
		conn.Write([]byte(msg + "\n"))
		if msg == "quit" {
			break
		}
	}
	fmt.Println("> echo-server is de-activated")
}
