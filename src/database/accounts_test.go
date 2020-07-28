package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

const (
	DRIVER           = "mysql"
	DATA_SOURCE_NAME = "root:toor@tcp([localhost]:3306)/panel?parseTime=true" // TODO decide database name
)

var db DB

func init() {
	var initErr error
	d, initErr := sql.Open(DRIVER, DATA_SOURCE_NAME)
	if initErr != nil {
		panic(initErr)
	}
	db.Conn = d
}

func TestGetAccounts(t *testing.T) {
	accounts, err := db.GetAccounts()
	if err != nil {
		t.Fatal(err)
	}
	wanted := Account{
		Id:       1,
		Name:     "user1",
		Expire:   nil,
		Os:       "mac",
		Password: "",
	}
	if len(accounts) != 2 {
		t.Fatalf("failed to get data: wanted %d: actual: %d", 2, len(accounts))
	}
	if accounts[0].Name != wanted.Name {
		t.Errorf("failed to get name: actual: %v", accounts[0].Name)
	}
}
