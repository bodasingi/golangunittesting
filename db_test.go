// main_test.go
package main

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type MockDB struct {
	db *gorm.DB
}

func (m *MockDB) First(dest interface{}, conds ...interface{}) *gorm.DB {
	return m.db.First(dest, conds...)
}

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

	// Set up the mock expectations
	userID := uint(1)
	mock.ExpectQuery(`SELECT \* FROM "users" WHERE "users"."id" = \$1`).WithArgs(userID).WillReturnRows(
		sqlmock.NewRows([]string{"id", "name", "age"}).AddRow(1, "John Doe", 30),
	)

	// Create a mock DB instance using the GORM DB
	mockDB := &MockDB{db: gormDB}

	// Call the function with the mock DB
	user, err := GetUserByID(mockDB, userID)

	// Verify the result
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if user.ID != 1 || user.Name != "John Doe" || user.Age != 30 {
		t.Errorf("Unexpected user data. Got: %+v", user)
	}

	// Check if all expectations were met
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}
