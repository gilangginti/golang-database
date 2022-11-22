package golangdatabase

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestDatabaseConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/golang_database")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
