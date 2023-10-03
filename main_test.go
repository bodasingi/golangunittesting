// main_test.go
package main

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestGetUserByID(t *testing.T) {
	// Create a mock DB
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error occurred: %s", err)
	}

	// Create a GORM DB instance using the mock DB
	gormDB, err := gorm.Open(sqlite.New(sqlite.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("An error occurred: %s", err)
	}

	// Close the mock DB when the test is done
	defer db.Close()

	// Mock the expected query
	userID := uint(1)
	mock.ExpectQuery(`SELECT \* FROM "users" WHERE "users"."id" = \$1`).WithArgs(userID).WillReturnRows(
		sqlmock.NewRows([]string{"id", "name", "age"}).AddRow(1, "John Doe", 30),
	)

	// Call the function that depends on the GORM DB
	user, err := GetUserByID(gormDB, userID)

	// Check if an error occurred
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	// Verify the result
	if user.ID != 1 || user.Name != "John Doe" || user.Age != 30 {
		t.Errorf("Unexpected user data. Got: %+v", user)
	}

	// Check if all expectations were met
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}
