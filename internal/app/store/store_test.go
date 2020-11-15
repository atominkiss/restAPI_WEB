package store_test

import (
	"os"
	"testing"
)

var databaseURL string

func TestMain(m *testing.M)  {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=172.17.0.3 dbname=restapi_test sslmode=disable user=admin password=Defacktos5"
	}

	os.Exit(m.Run())
}