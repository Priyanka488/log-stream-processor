package tcp

import (
	"fmt"
	"net"
	"sync"

	"github.com/Priyanka488/log-stream-processor/config"
)

func Init(wg *sync.WaitGroup) {

	defer wg.Done()

	var wg_tcp sync.WaitGroup
	// 1. create a listener
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", config.TCP_PORT))
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}

	// 3. keep listening for new connections
	for {
		// 2. start accepting connections
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting:", err.Error())
			return
		}
		wg_tcp.Add(1)
		go handleRequest(conn, &wg_tcp)
	}

}

// 4. keep listening on the same connection
func handleRequest(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		chunk := make([]byte, config.TCP_MESSAGE_SIZE)
		readBytes, err := conn.Read(chunk)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			return
		}
		fmt.Printf("Received data: %v", string(chunk[:readBytes]))
	}
}
