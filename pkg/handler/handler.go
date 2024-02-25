package handler

import (
	"context"
	"fmt"
	"sync"

	"github.com/Priyanka488/log-stream-processor/config"
	"github.com/Priyanka488/log-stream-processor/pkg/models"
)

func processEvent(i int, ch chan models.Event, wg *sync.WaitGroup, ctx context.Context) {

	defer wg.Done()
	defer fmt.Println("Handler ", i, " is done")

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Closing handler: ", i)
			return
		case event, ok := <-ch:
			if !ok {
				fmt.Println("Channel closed")
				return
			}
			event.Display()
			// process event

			// filter event
		}
	}

}

func Init(ch chan models.Event, wg *sync.WaitGroup, ctx context.Context) {
	for i := 0; i < config.MAX_HANDLERS; i++ {
		wg.Add(1)
		go processEvent(i+1, ch, wg, ctx)
	}
}
