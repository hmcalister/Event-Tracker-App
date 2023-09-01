package database

import (
	"os"
	"testing"

	"gorm.io/gorm"
)

const TEST_DATABASE_FILE string = "testDatabase.db"

func TestCreateDatabaseConnectionFailure(t *testing.T) {
	var err error
	t.Cleanup(func() {
		os.Chmod(TEST_DATABASE_FILE, 0664)
		os.Remove(TEST_DATABASE_FILE)
	})

	os.Create(TEST_DATABASE_FILE)
	os.Chmod(TEST_DATABASE_FILE, 0000)
	if _, err = createDatabaseConnection(TEST_DATABASE_FILE); err == nil {
		t.Errorf("Creating a database connection to a file with 000 permission succeeded - failure expected.")
	}
}

func TestCreateDatabaseConnectionAndSetup(t *testing.T) {
	var databaseConnection *gorm.DB
	var err error

	t.Cleanup(func() { os.Remove(TEST_DATABASE_FILE) })

	if databaseConnection, err = createDatabaseConnection(TEST_DATABASE_FILE); err != nil {
		t.Errorf("error during database connection creation: %v", err)
	}

	if err = databaseSetup(databaseConnection); err != nil {
		t.Errorf("error during database connection setup: %v", err)
	}
}

func TestCreateDatabase(t *testing.T) {
	var err error

	t.Cleanup(func() { os.Remove(TEST_DATABASE_FILE) })
	if _, err = CreateDatabase(TEST_DATABASE_FILE); err != nil {
		t.Errorf("error during database connection: %v", err)
	}
}

func TestCreateDatabaseWithFilePermissionError(t *testing.T) {
	var err error
	t.Cleanup(func() {
		os.Chmod(TEST_DATABASE_FILE, 0664)
		os.Remove(TEST_DATABASE_FILE)
	})

	os.Create(TEST_DATABASE_FILE)
	os.Chmod(TEST_DATABASE_FILE, 0000)
	if _, err = CreateDatabase(TEST_DATABASE_FILE); err == nil {
		t.Errorf("Creating a database connection to a file with 000 permission succeeded - failure expected.")
	}
}
