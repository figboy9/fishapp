package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetUser(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	db.Close()
	t.Log("success")
}
