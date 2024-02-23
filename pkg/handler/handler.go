package handler

import (
	"context"
	"fmt"
	"sync"

	"github.com/Priyanka488/log-stream-processor/config"
	"github.com/Priyanka488/log-stream-processor/pkg/filter"
	"github.com/Priyanka488/log-stream-processor/pkg/models"
	"github.com/Priyanka488/log-stream-processor/pkg/processor"
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
			processedEvent := processor.ProcessEvent(event)
			// Example filter: Only process INFO severity events.
			filter := filter.SeverityFilter("INFO")
			if filter(processedEvent) {
				fmt.Printf("Handler %d: %v\n", i, processedEvent)
				// Here you would forward the event to the egress layer.
			}
		}
	}
}

func Init(ch chan models.Event, wg *sync.WaitGroup, ctx context.Context) {
	for i := 0; i < config.MAX_HANDLERS; i++ {
		wg.Add(1)
		go processEvent(i+1, ch, wg, ctx)
	}
}
