package processor

import (
	"github.com/Priyanka488/log-stream-processor/pkg/models"
)

// ProcessEvent modifies or enriches an event.
func ProcessEvent(event models.Event) models.Event {
	// Example: Adding a timestamp or modifying the event's body.
	// This is a placeholder for your processing logic.

	if event != nil {
		if systemLog, ok := event.(*models.SystemLog); ok {
			systemLog.Log.SetTimeCreated()
		}
	}

	switch e := event.(type) {
	case models.SystemLog:
		e.Body += " - Processed"
		return e
	default:
		// If the event type is unknown, return it unmodified.
		return event
	}
}
