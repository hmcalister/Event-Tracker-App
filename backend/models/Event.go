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

