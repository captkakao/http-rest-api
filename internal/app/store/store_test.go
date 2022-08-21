package store_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=localhost user=test_user password=qweqwe123 dbname=test_db port=5432 sslmode=disable"
	}

	os.Exit(m.Run())
}
