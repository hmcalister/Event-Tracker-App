package main

import (
	"context"
	"fmt"
	"hmcalister/EventTrackerApp/backend/database"
	"hmcalister/EventTrackerApp/backend/models"
	"log"

	"gorm.io/gorm"
)

// App struct
type App struct {
	ctx                context.Context
	databaseConnection *gorm.DB
	allEventsList      []*models.Event
	allEventsMap       map[uint]*models.Event
}

// InitApp creates a new App application struct
func InitApp() *App {
	databaseConnection, err := database.CreateDatabase(database.DEFAULT_DATABASE_FILE)
	if err != nil {
		log.Fatalf("error during creation of database: %v\n", err)
	}

	app := &App{
		databaseConnection: databaseConnection,
	}

	app.allEventsList, app.allEventsMap = app.initAllEventsFromDatabase()
	return app
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) initAllEventsFromDatabase() ([]*models.Event, map[uint]*models.Event) {
	var allEvents []*models.Event
	a.databaseConnection.Find(&allEvents)

	allEventsMap := make(map[uint]*models.Event)
	for _, event := range allEvents {
		allEventsMap[event.ID] = event
	}
	return allEvents, allEventsMap
}

func (a *App) GetAllEvents() []*models.Event {
	var allEvents []*models.Event
	a.databaseConnection.Find(&allEvents)
	return allEvents
}

func (a *App) CreateEvent(newEvent *models.Event) *models.Event {
	models.RegisterNewEventInDatabase(newEvent, a.databaseConnection)
	return newEvent
}

func (a *App) UpdateEvent(updatedEvent *models.Event) *models.Event {
	updatedEvent.UpdateEventInDatabase(a.databaseConnection)
	return updatedEvent
}

func (a *App) DismissEvent(event *models.Event) *models.Event {
	dismissedEvent := event.DismissEvent(a.databaseConnection)
	return dismissedEvent
}

func (a *App) DeleteEvent(event *models.Event) *models.Event {
	event.IsRecurring = false
	return a.DismissEvent(event)
}

func (a *App) PrintEventStruct(event *models.Event) {
	fmt.Printf("%+v\n", event)
}
