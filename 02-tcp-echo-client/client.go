package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Println("> echo-client is activated")
	conn, _ := net.Dial("tcp", "127.0.0.1:65456")
	defer conn.Close()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		conn.Write([]byte(msg + "\n"))
		response, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("> received:", strings.TrimSpace(response))
		if msg == "quit" {
			break
		}
	}
	fmt.Println("> echo-client is de-activated")
}
