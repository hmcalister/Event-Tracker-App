package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Name            string
	StartTime       time.Time
	TimeoutDuration time.Duration
	IsRecurring     bool
}

// Create a new Event
func RegisterNewEventInDatabase(newEvent *Event, databaseConnection *gorm.DB) error {
	result := databaseConnection.Create(newEvent)
	return result.Error
}

// Update a specific event
func (event *Event) UpdateEventInDatabase(databaseConnection *gorm.DB) error {
	result := databaseConnection.Save(&event)
	return result.Error
}

// Delete an event
func (event *Event) DeleteEventInDatabase(databaseConnection *gorm.DB) error {
	result := databaseConnection.Delete(&event)
	return result.Error
}

// Get the time until the event trigger
func (event *Event) TriggerTime() time.Time {
	return event.StartTime.Add(event.TimeoutDuration)
}

// Check if the event is currently triggered
func (event *Event) IsTriggered() bool {
	return time.Now().After(event.TriggerTime())
}

// Set the event start time.
// If the database connection passed is not nil the struct is also updated in the database
func (event *Event) SetStartTime(newStartTime time.Time, databaseConnection *gorm.DB) {
	event.StartTime = newStartTime
	if databaseConnection != nil {
		event.UpdateEventInDatabase(databaseConnection)
	}
}

// Set the event TimeoutDuration
func (event *Event) SetTimeoutDuration(newTimeoutDuration time.Duration, databaseConnection *gorm.DB) {
	event.TimeoutDuration = newTimeoutDuration
	if databaseConnection != nil {
		event.UpdateEventInDatabase(databaseConnection)
	}
}

// Reset the event by setting the start time to time.Now().
//
// Returns the event or nil. Event is returned in the case of a recurring event. For a non-recurring event, nil is returned.
// This allows for dismissed non-recurring events to be removed appropriately, keeping consistency with the database.
func (event *Event) DismissEvent(databaseConnection *gorm.DB) *Event {
	if event.IsRecurring {
		event.StartTime = time.Now()
		if databaseConnection != nil {
			event.UpdateEventInDatabase(databaseConnection)
		}
		return event
	} else {
		if databaseConnection != nil {
			event.DeleteEventInDatabase(databaseConnection)
		}
		return nil
	}
}
