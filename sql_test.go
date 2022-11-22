package golangdatabase

import (
	"context"
	"fmt"
	"testing"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	insert := "INSERT INTO customer(id, name) VALUES('1', 'Gilang')"
	_, err := db.ExecContext(ctx, insert)

	if err != nil {
		panic(err)
	}

	fmt.Println("SUCCESS INSERT")
}
