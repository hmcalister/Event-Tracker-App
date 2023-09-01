package models_test

import (
	"hmcalister/EventTrackerApp/backend/database"
	"hmcalister/EventTrackerApp/backend/models"
	"os"
	"testing"
	"time"
)

const TEST_DATABASE_FILE string = "testDatabase.db"

var databaseConnection, _ = database.CreateDatabase(TEST_DATABASE_FILE)

// We use a main test here so we can cleanup the database file later
func TestAllWithDatabase(t *testing.T) {
	t.Cleanup(func() { os.Remove(TEST_DATABASE_FILE) })

	t.Run("TestCreateEventsInDatabase", _TestCreateEventsInDatabase)
	t.Run("TestUpdateEventsInDatabase", _TestUpdateEventsInDatabase)
	t.Run("TestDeleteEventsInDatabase", _TestDeleteEventsInDatabase)
	t.Run("TestSetStartTimeWithDatabase", _TestSetStartTimeWithDatabase)
	t.Run("TestSetTimeoutDurationWithDatabase", _TestSetTimeoutDurationWithDatabase)
	t.Run("TestDismissEventWithDatabase", _TestDismissEventWithDatabase)

}

func _TestCreateEventsInDatabase(t *testing.T) {
	var err error
	t.Cleanup(func() { databaseConnection.Migrator().DropTable(&models.Event{}) })

	databaseConnection.AutoMigrate(&models.Event{})
	for testEventIndex, testEvent := range testEventStructs {
		err = models.RegisterNewEventInDatabase(testEvent, databaseConnection)
		if err != nil {
			t.Errorf("failed to register event in database for Event struct %v (index %v)", testEvent, testEventIndex)
		}
	}
}

func _TestUpdateEventsInDatabase(t *testing.T) {
	var err error
	t.Cleanup(func() { databaseConnection.Migrator().DropTable(&models.Event{}) })
	databaseConnection.AutoMigrate(&models.Event{})

	for testEventIndex, testEvent := range testEventStructs {
		oldTime := testEvent.StartTime
		newTime := time.Now().Add(time.Hour * time.Duration(testEventIndex+1))

		models.RegisterNewEventInDatabase(testEvent, databaseConnection)

		testEvent.StartTime = newTime
		if err = testEvent.UpdateEventInDatabase(databaseConnection); err != nil {
			t.Errorf("error during updating event in database for Event struct %v (index %v)", testEvent, testEventIndex)
		}

		testEvent.StartTime = oldTime
		if err = testEvent.UpdateEventInDatabase(databaseConnection); err != nil {
			t.Errorf("error during updating event in database for Event struct %v (index %v)", testEvent, testEventIndex)
		}
	}
}

func _TestDeleteEventsInDatabase(t *testing.T) {
	var err error
	t.Cleanup(func() { databaseConnection.Migrator().DropTable(&models.Event{}) })
	databaseConnection.AutoMigrate(&models.Event{})

	for testEventIndex, testEvent := range testEventStructs {
		models.RegisterNewEventInDatabase(testEvent, databaseConnection)

		if err = testEvent.DeleteEventInDatabase(databaseConnection); err != nil {
			t.Errorf("error during deleting event in database for Event struct %v (index %v)", testEvent, testEventIndex)
		}
	}
}

func _TestSetStartTimeWithDatabase(t *testing.T) {
	t.Cleanup(func() { databaseConnection.Migrator().DropTable(&models.Event{}) })
	databaseConnection.AutoMigrate(&models.Event{})

	for testEventIndex, testEvent := range testEventStructs {
		oldTime := testEvent.StartTime
		newTime := time.Now().Add(time.Hour * time.Duration(testEventIndex+1))

		testEvent.SetStartTime(newTime, databaseConnection)
		if testEvent.StartTime != newTime {
			t.Errorf("expected StartTime (%#v) does not match found StartTime (%#v) for Event struct (index %v)", newTime, oldTime, testEventIndex)
		}

		// Reset struct and test again
		testEvent.SetStartTime(oldTime, databaseConnection)
		if testEvent.StartTime != oldTime {
			t.Errorf("expected StartTime (%#v) does not match found StartTime (%#v) for Event struct (index %v) (during undo)", newTime, oldTime, testEventIndex)
		}
	}
}

func _TestSetTimeoutDurationWithDatabase(t *testing.T) {
	t.Cleanup(func() { databaseConnection.Migrator().DropTable(&models.Event{}) })
	databaseConnection.AutoMigrate(&models.Event{})

	for testEventIndex, testEvent := range testEventStructs {
		oldTimeoutDuration := testEvent.TimeoutDuration
		newTimeoutDuration := time.Hour * time.Duration(testEventIndex+1)

		testEvent.SetTimeoutDuration(newTimeoutDuration, databaseConnection)
		if testEvent.TimeoutDuration != newTimeoutDuration {
			t.Errorf("expected TimeoutDuration (%#v) does not match found TimeoutDuration (%#v) for Event struct (index %v)", newTimeoutDuration, oldTimeoutDuration, testEventIndex)
		}

		// Reset struct and test again
		testEvent.SetTimeoutDuration(oldTimeoutDuration, databaseConnection)
		if testEvent.TimeoutDuration != oldTimeoutDuration {
			t.Errorf("expected TimeoutDuration (%#v) does not match found TimeoutDuration (%#v) for Event struct (index %v) (during undo)", newTimeoutDuration, oldTimeoutDuration, testEventIndex)
		}
	}
}

func _TestDismissEventWithDatabase(t *testing.T) {
	t.Cleanup(func() { databaseConnection.Migrator().DropTable(&models.Event{}) })
	databaseConnection.AutoMigrate(&models.Event{})

	for testEventIndex, testEvent := range testEventStructs {
		oldTime := testEvent.StartTime

		expectedReturnsNil := dismissEventReturnsNil[testEventIndex]
		foundReturnsNil := testEvent.DismissEvent(databaseConnection) == nil
		if expectedReturnsNil != foundReturnsNil {
			t.Errorf("expected returnsNil (%#v) does not match found returnsNil (%#v) for Event struct %v (index %v)", expectedReturnsNil, foundReturnsNil, testEvent, testEventIndex)
		}

		testEvent.StartTime = oldTime
	}
}

// TODO: Add tests to check if events are actually updated in database (i.e. query database, compare found structs to expected ones)
