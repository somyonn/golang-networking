package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

func main() {
	fmt.Println("> echo-client is activated")
	conn, err := net.Dial("tcp", "127.0.0.1:65456")
	if err != nil {
		fmt.Println("> connect() failed and program terminated")
		return
	}
	defer conn.Close()

	var wg sync.WaitGroup
	wg.Add(2)

	// 송신 고루틴
	go func() {
		defer wg.Done()
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("> ")
			msg, _ := reader.ReadString('\n')
			msg = strings.TrimSpace(msg)
			conn.Write([]byte(msg + "\n"))
			if msg == "quit" {
				return
			}
		}
	}()

	// 수신 고루틴
	go func() {
		defer wg.Done()
		for {
			response, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				return
			}
			fmt.Println("> received:", strings.TrimSpace(response))
		}
	}()

	wg.Wait()
	fmt.Println("> echo-client is de-activated")
}
