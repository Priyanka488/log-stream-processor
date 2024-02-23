package filter

import (
	"github.com/Priyanka488/log-stream-processor/pkg/models"
)

// FilterFunc defines the signature for filter functions.
type FilterFunc func(models.Event) bool

// SeverityFilter returns a filter function that filters events based on severity.
func SeverityFilter(severity string) FilterFunc {
	return func(event models.Event) bool {
		if e, ok := event.(models.SystemLog); ok {
			return e.Severity == severity
		}
		return false
	}
}

// FilterEvents filters events based on a given filter function.
func FilterEvents(events []models.Event, filter FilterFunc) []models.Event {
	var filteredEvents []models.Event
	for _, event := range events {
		if filter(event) {
			filteredEvents = append(filteredEvents, event)
		}
	}
	return filteredEvents
}
