package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

var clientID int
var clients = make([]net.Conn, 0)
var mutex sync.Mutex

func main() {
	fmt.Println("> echo-server is activated")
	listener, err := net.Listen("tcp", "127.0.0.1:65456")
	if err != nil {
		fmt.Println("> bind() failed and program terminated")
		return
	}
	defer listener.Close()

	mutex.Lock()
	clientID++
	fmt.Println("> server loop running in thread (main thread): Thread-", clientID)
	mutex.Unlock()

	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			input, _ := reader.ReadString('\n')
			if strings.TrimSpace(input) == "quit" {
				remain()
			}
		}
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("> listen() failed and program terminated")
			return
		}

		mutex.Lock()
		clientID++
		clients = append(clients, conn)
		mutex.Unlock()

		go handleConnection(conn, clientID)

	}
}

func handleConnection(conn net.Conn, clientID int) {
	defer conn.Close()
	defer removeClient(conn)
	clientAddress := conn.RemoteAddr().(*net.TCPAddr)
	fmt.Printf("> client connected by IP address %s with Port number %d, Thread-%d\n", clientAddress.IP.String(), clientAddress.Port, clientID)
	reader := bufio.NewReader(conn)
	for {
		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		fmt.Printf("> received ( %s ) and echoed to  %d clients\n", msg, len(clients))

		mutex.Lock()
		for _, conn := range clients {
			conn.Write([]byte(msg + "\n"))
		}
		mutex.Unlock()
		
		if msg == "quit" {
			break
		}
	}
}

func removeClient(conn net.Conn) {
	mutex.Lock()
	defer mutex.Unlock()
	for i, c := range clients {
		if c == conn {
			clients = append(clients[:i], clients[i+1:]...)
			break
		}
	}
}

func remain() {
	mutex.Lock()
	defer mutex.Unlock()
	clientID--
	if clientID == 0 {
		fmt.Println("stop procedure started")
		fmt.Println("echo-server is de-activated")
		os.Exit(0)
	}

	fmt.Printf("active threads are remained : %d threads\n", clientID)
}

func broadcast(msg string) {
	
}
