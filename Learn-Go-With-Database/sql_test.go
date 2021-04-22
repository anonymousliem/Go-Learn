package Learn_Go_With_Database

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func TestExecSQL(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO customer(id,name) VALUES ('mus','mus')"

	_, err := db.ExecContext(ctx, script)

	if err != nil{
		panic(err)
	}

	fmt.Println("Success Insert Data")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "SELECT id, name FROM customer"

	rows, err := db.QueryContext(ctx, script)

	if err != nil{
		panic(err)
	}

	defer rows.Close()

	for rows.Next(){
		var id, name string
		err = rows.Scan(&id, &name)
		if err != nil{
			panic(err)
		}

		fmt.Println("ID", id)
		fmt.Println("Name", name)
	}
}


func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "SELECT id, name, email, ballance, rating, created_at, birth_date, married FROM customer"

	rows, err := db.QueryContext(ctx, script)

	if err != nil{
		panic(err)
	}

	defer rows.Close()

	for rows.Next(){
		var id, name string
		var email sql.NullString
		var ballance int32
		var rating float64
		var birthDate sql.NullTime
		var createAt time.Time
		var married bool

		err = rows.Scan(&id, &name, &email, &ballance, &rating, &birthDate, &createAt, &married)
		if err != nil{
			panic(err)
		}

		fmt.Println("ID", id)
		fmt.Println("Name", name)
		if email.Valid{
			fmt.Println("email", email.String)
		}

		fmt.Println("balance", ballance)
		fmt.Println("rating", rating)
		fmt.Println("birth date", birthDate.Time)
		fmt.Println("create at", createAt)
		fmt.Println("married", married)
	}
}