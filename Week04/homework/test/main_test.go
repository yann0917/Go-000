package test

import (
	. "github.com/wlxpkg/base"
	"database/sql"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

var db *sql.DB
var Mock sqlmock.Sqlmock
var err error

func setup() {
	var t *testing.T
	db, Mock, err = sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	DB, _ = gorm.Open("mysql", db)
}

func shutdown() {
	defer db.Close()
}
