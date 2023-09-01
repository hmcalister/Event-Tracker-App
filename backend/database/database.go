package database

import (
	"hmcalister/EventTrackerApp/backend/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const DEFAULT_DATABASE_FILE string = "database.db"

// Setup the database by automigrating the models defined in hmcalister/EventTrackerApp/backend/models.
func databaseSetup(dbConn *gorm.DB) error {
	if err := dbConn.AutoMigrate(&models.Event{}); err != nil {
		return err
	}

	return nil
}

// create a database file and return the connection
func createDatabaseConnection(databaseFile string) (*gorm.DB, error) {
	dbConn, err := gorm.Open(sqlite.Open(databaseFile), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return dbConn, nil
}

// Create a database connection by opening an sqlite database at a specific filepath.
// The database, once created, is setup using databaseSetup().
func CreateDatabase(databaseFile string) (*gorm.DB, error) {
	dbConn, err := createDatabaseConnection(databaseFile)
	if err != nil {
		return nil, err
	}

	err = databaseSetup(dbConn)
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}
