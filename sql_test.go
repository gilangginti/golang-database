package golangdatabase

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	insert := "INSERT INTO customer(id, name, email, balance, rating, birth_date, married) VALUES('1', 'Gilang', 'gilang@gmail.com', 10000, 90.0, '1999-04-07', true)"
	_, err := db.ExecContext(ctx, insert)

	if err != nil {
		panic(err)
	}

	fmt.Println("SUCCESS INSERT")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	getAllCustomer := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, getAllCustomer)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birthDate sql.NullTime
		var createdAt time.Time
		var married bool

		err = rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}

		fmt.Println("=======")
		fmt.Println("id : ", id)
		fmt.Println("name : ", name)
		if email.Valid {
			fmt.Println("email : ", email)
		}
		fmt.Println("balance : ", balance)
		fmt.Println("rating : ", rating)

		if birthDate.Valid {
			fmt.Println("birthDate : ", birthDate)
		}

		fmt.Println("married : ", married)
		fmt.Println("createdAt : ", createdAt)
		fmt.Println("=======")
	}
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin; #"
	password := "admin"

	getUser := "SELECT username, password FROM user WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1"
	rows, err := db.QueryContext(ctx, getUser)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		var password string
		err := rows.Scan(&username, &password)
		if err != nil {
			panic(err)
		}
		fmt.Println("SUKSES LOGIN")
	} else {
		fmt.Println("GAGAL LOGIN")
	}
}

func TestSqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin; #"
	password := "admin"

	getUser := "SELECT username, password FROM user WHERE username = ? AND password = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, getUser, username, password)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		var password string
		err := rows.Scan(&username, &password)
		if err != nil {
			panic(err)
		}
		fmt.Println("SUKSES LOGIN")
	} else {
		fmt.Println("GAGAL LOGIN")
	}
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "gilang@gmail.com"
	comment := "wakwaw"
	insert := "INSERT INTO comments(email, comment) VALUES(?,?)"
	result, err := db.ExecContext(ctx, insert, email, comment)

	if err != nil {
		panic(err)
	}
	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("SUCCESS INSERT WITH LAST ID", insertId)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	insert := "INSERT INTO comments(email, comment) VALUES(?,?)"
	statement, err := db.PrepareContext(ctx, insert)
	defer statement.Close()
	if err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++ {
		email := "gilang" + strconv.Itoa(i) + "@gmail.com"
		comment := "Komentar Ke " + strconv.Itoa(i)
		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}
		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Id Ke", id)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	insert := "INSERT INTO comments(email, comment) VALUES(?,?)"

	for i := 0; i < 10; i++ {
		email := "gilang" + strconv.Itoa(i) + "@gmail.com"
		comment := "Komentar Ke " + strconv.Itoa(i)
		result, err := tx.ExecContext(ctx, insert, email, comment)
		if err != nil {
			panic(err)
		}
		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Id Ke", id)
	}
	//err = tx.Commit()
	err = tx.Rollback()
	if err != nil {
		panic(err)
	}
}
