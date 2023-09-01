package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
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

