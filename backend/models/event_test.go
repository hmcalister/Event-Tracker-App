package models_test

import (
	"hmcalister/EventTrackerApp/backend/models"
	"testing"
	"time"
)

var TIME_NOW time.Time = time.Now()

var testEventStructs = []*models.Event{
	// Now with 10 Second timeout
	{
		Name:            "Now with 10 Second timeout",
		StartTime:       TIME_NOW,
		TimeoutDuration: time.Second * 10,
	},

	// Now with 10 Hour timeout
	{
		Name:            "Now with 10 Hour timeout",
		StartTime:       TIME_NOW,
		TimeoutDuration: time.Hour * 10,
	},

	// Now (minus 1 Hour) with 10 Second timeout
	{
		Name:            "Now (minus 1 Hour) with 10 Second timeout",
		StartTime:       TIME_NOW.Add(-time.Hour),
		TimeoutDuration: time.Second * 10,
	},

	// Now (minus 1 Hour) with 10 Hour timeout
	{
		Name:            "Now (minus 1 Hour) with 10 Hour timeout",
		StartTime:       TIME_NOW.Add(-time.Hour),
		TimeoutDuration: time.Hour * 10,
	},

	// UnixEpoch with 10 Second timeout
	{
		Name:            "UnixEpoch with 10 Second timeout",
		StartTime:       time.UnixMicro(0),
		TimeoutDuration: time.Second * 10,
	},

	// UnixEpoch with 10 Hour timeout
	{
		Name:            "UnixEpoch with 10 Hour timeout",
		StartTime:       time.UnixMicro(0),
		TimeoutDuration: time.Hour * 10,
	},

	// UnixEpoch with 10 Second timeout and recurring
	{
		Name:            "UnixEpoch with 10 Second timeout and recurring",
		StartTime:       time.UnixMicro(0),
		TimeoutDuration: time.Second * 10,
		IsRecurring:     true,
	},

	// UnixEpoch with 10 Hour timeout and recurring
	{
		Name:            "UnixEpoch with 10 Hour timeout and recurring",
		StartTime:       time.UnixMicro(0),
		TimeoutDuration: time.Hour * 10,
		IsRecurring:     true,
	},
}

var expectedTriggerTimes = []time.Time{
	TIME_NOW.Add(time.Second * 10),           //Now with 10 Second timeout
	TIME_NOW.Add(time.Hour * 10),             //Now with 10 Hour timeout
	TIME_NOW.Add(time.Second*10 - time.Hour), //Now (minus 1 Hour) with 10 Second timeout
	TIME_NOW.Add(time.Hour*10 - time.Hour),   //Now (minus 1 Hour) with 10 Hour timeout
	time.Unix(0, 0).Add(time.Second * 10),    //UnixEpoch with 10 Second timeout
	time.Unix(0, 0).Add(time.Hour * 10),      //UnixEpoch with 10 Hour timeout
	time.Unix(0, 0).Add(time.Second * 10),    // UnixEpoch with 10 Second timeout and recurring
	time.Unix(0, 0).Add(time.Hour * 10),      // UnixEpoch with 10 Hour timeout and recurring
}

func TestTriggerTime(t *testing.T) {
	for testEventIndex, testEvent := range testEventStructs {
		expectedTriggerTime := expectedTriggerTimes[testEventIndex]
		foundTriggerTime := testEvent.TriggerTime()
		if expectedTriggerTime != foundTriggerTime {
			t.Errorf("expected trigger time (%#v) does not match found trigger time (%#v) for Event struct %+v (index %v)", expectedTriggerTime, foundTriggerTime, testEvent, testEventIndex)
		}
	}
}

var expectedIsTriggeredValues = []bool{
	false, //Now with 10 Second timeout
	false, //Now with 10 Hour timeout
	true,  //Now (minus 1 Hour) with 10 Second timeout
	false, //Now (minus 1 Hour) with 10 Hour timeout
	true,  //UnixEpoch with 10 Second timeout
	true,  //UnixEpoch with 10 Hour timeout
	true,  // UnixEpoch with 10 Second timeout and recurring
	true,  // UnixEpoch with 10 Hour timeout and recurring
}

func TestIsTriggered(t *testing.T) {
	for testEventIndex, testEvent := range testEventStructs {
		expectedIsTriggered := expectedIsTriggeredValues[testEventIndex]
		foundIsTriggered := testEvent.IsTriggered()
		if expectedIsTriggered != foundIsTriggered {
			t.Errorf("expected isTriggered (%#v) does not match found isTriggered (%#v) for Event struct %+v (index %v)", expectedIsTriggered, foundIsTriggered, testEvent, testEventIndex)
		}
	}
}

func TestSetStartTime(t *testing.T) {
	for testEventIndex, testEvent := range testEventStructs {
		oldTime := testEvent.StartTime
		newTime := time.Now().Add(time.Hour * time.Duration(testEventIndex+1))

		testEvent.SetStartTime(newTime, nil)
		if testEvent.StartTime != newTime {
			t.Errorf("expected StartTime (%#v) does not match found StartTime (%#v) for Event struct (index %v)", newTime, oldTime, testEventIndex)
		}

		// Reset struct and test again
		testEvent.SetStartTime(oldTime, nil)
		if testEvent.StartTime != oldTime {
			t.Errorf("expected StartTime (%#v) does not match found StartTime (%#v) for Event struct (index %v) (during undo)", newTime, oldTime, testEventIndex)
		}
	}
}

func TestSetTimeoutDuration(t *testing.T) {
	for testEventIndex, testEvent := range testEventStructs {
		oldTimeoutDuration := testEvent.TimeoutDuration
		newTimeoutDuration := time.Hour * time.Duration(testEventIndex+1)

		testEvent.SetTimeoutDuration(newTimeoutDuration, nil)
		if testEvent.TimeoutDuration != newTimeoutDuration {
			t.Errorf("expected TimeoutDuration (%#v) does not match found TimeoutDuration (%#v) for Event struct (index %v)", newTimeoutDuration, oldTimeoutDuration, testEventIndex)
		}

		// Reset struct and test again
		testEvent.SetTimeoutDuration(oldTimeoutDuration, nil)
		if testEvent.TimeoutDuration != oldTimeoutDuration {
			t.Errorf("expected TimeoutDuration (%#v) does not match found TimeoutDuration (%#v) for Event struct (index %v) (during undo)", newTimeoutDuration, oldTimeoutDuration, testEventIndex)
		}
	}
}

var dismissEventReturnsNil = []bool{
	true,  //Now with 10 Second timeout
	true,  //Now with 10 Hour timeout
	true,  //Now (minus 1 Hour) with 10 Second timeout
	true,  //Now (minus 1 Hour) with 10 Hour timeout
	true,  //UnixEpoch with 10 Second timeout
	true,  //UnixEpoch with 10 Hour timeout
	false, // UnixEpoch with 10 Second timeout and recurring
	false, // UnixEpoch with 10 Hour timeout and recurring
}

func TestDismissEvent(t *testing.T) {
	for testEventIndex, testEvent := range testEventStructs {
		oldTime := testEvent.StartTime
		expectedReturnsNil := dismissEventReturnsNil[testEventIndex]
		foundReturnsNil := testEvent.DismissEvent(nil) == nil
		if expectedReturnsNil != foundReturnsNil {
			t.Errorf("expected returnsNil (%#v) does not match found returnsNil (%#v) for Event struct %v (index %v)", expectedReturnsNil, foundReturnsNil, testEvent, testEventIndex)
		}
		testEvent.StartTime = oldTime
	}
}
