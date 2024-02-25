package tcp

import (
	"context"
	"fmt"
	"net"
	"strings"
	"sync"

	"github.com/Priyanka488/log-stream-processor/config"
	"github.com/Priyanka488/log-stream-processor/pkg/models"
	"github.com/Priyanka488/log-stream-processor/pkg/processor"
)

// 5. Add Context Cancellation to TCP Server

func Init(wg *sync.WaitGroup, ctx context.Context, ch chan models.Event) {

	defer wg.Done()

	var wg_tcp sync.WaitGroup
	// 1. create a listener
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", config.TCP_PORT))
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}

	// 3. keep listening for new connections
	go func() {
		// 2. start accepting connections
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("Error accepting:", err.Error())
				return
			}
			wg_tcp.Add(1)
			go handleRequest(conn, &wg_tcp, ctx, ch)
		}
	}()

	<-ctx.Done()
	fmt.Println("Closing TCP server, context cancelled")
	if err := listener.Close(); err != nil {
		fmt.Println("Error closing listener:", err.Error())
	}
	wg_tcp.Wait()
	fmt.Println("TCP server closed")
}

// 4. keep listening on the same connection
func handleRequest(conn net.Conn, wg *sync.WaitGroup, ctx context.Context, ch chan models.Event) {
	defer wg.Done()
	fmt.Println("New connection")
	connect_ch := make(chan bool)
	go func(connect_ch chan bool) {
		for {
			chunk := make([]byte, config.TCP_MESSAGE_SIZE)
			readBytes, err := conn.Read(chunk)
			if err != nil {
				fmt.Println("Error reading:", err.Error())
				connect_ch <- true
				return
			}
			// 6. Send data to handler
			data_str := string(chunk[:readBytes])
			// trim null characters
			data_str = strings.TrimSpace(data_str)
			ch <- processor.ProcessStringToEvent(data_str)
			fmt.Printf("Received data: %v", data_str)
		}
	}(connect_ch)

	select {
	case <-connect_ch:
		fmt.Println("Closing connection, client disconnected")
	case <-ctx.Done():
		fmt.Println("Closing connection, context cancelled")
	}

	if err := conn.Close(); err != nil {
		fmt.Println("Error closing connection:", err.Error())
	}
}
