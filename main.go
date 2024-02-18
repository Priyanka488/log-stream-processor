package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/Priyanka488/log-stream-processor/pkg/handler"
	"github.com/Priyanka488/log-stream-processor/pkg/models"
)

func simulateIngress(ch chan models.Event) {
	for i := 0; i < 10; i++ {
		ch <- models.SystemLog{
			Log:      models.Log{ID: i, Source: "App", Body: "System is running"},
			Severity: "INFO",
		}
		time.Sleep(1 * time.Second)
	}
}

func listenForCancel(cancel context.CancelFunc, wg *sync.WaitGroup) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	fmt.Println("Received signal to cancel")
	cancel()
	wg.Done()
}

func main() {
	ch := make(chan models.Event, 100)
	var wg sync.WaitGroup
]
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go listenForCancel(cancel, &wg)

	simulateIngress(ch)
	close(ch)
	defer wg.Wait()
}
