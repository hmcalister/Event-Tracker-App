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
