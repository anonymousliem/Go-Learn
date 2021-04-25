package Learn_Go_With_Database

import (
	"context"
	"database/sql"
	"fmt"
	"learn-go-with-database/entity"
	"learn-go-with-database/repository"
	"strconv"
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

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "admin"

	script := "SELECT username FROM user WHERE username='"+username+"' AND password='"+password+"' LIMIT 1 "
	println(script)
	rows, err := db.QueryContext(ctx, script)

	if err != nil{
		panic(err)
	}

	defer rows.Close()

	if rows.Next(){
		var username string
		err := rows.Scan(&username)
		if err != nil{
			panic(err)
		}
		fmt.Println("sukses login", username)
	}else{
		fmt.Println("Gagal Login")
	}
}

func TestSqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "admin"

	script := "SELECT username FROM user WHERE username=? AND password=? LIMIT 1 "
	println(script)
	rows, err := db.QueryContext(ctx, script, username, password)

	if err != nil{
		panic(err)
	}

	defer rows.Close()

	if rows.Next(){
		var username string
		err := rows.Scan(&username)
		if err != nil{
			panic(err)
		}
		fmt.Println("sukses login", username)
	}else{
		fmt.Println("Gagal Login")
	}
}

func TestExecSQLSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	username := "awok"
	password := "awoko"
	ctx := context.Background()
	script := "INSERT INTO user(username,password) VALUES (?,?)"

	_, err := db.ExecContext(ctx, script, username, password)

	if err != nil{
		panic(err)
	}

	fmt.Println("Success Insert Data")
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	email := "awok@gmail.com"
	comment := "no comment"
	ctx := context.Background()
	script := "INSERT INTO comments(email,comment) VALUES (?,?)"

	result, err := db.ExecContext(ctx, script, email, comment)

	if err != nil{
		panic(err)
	}

	insertId, err := result.LastInsertId()

	if err != nil{
		panic(err)
	}
	fmt.Println("Success Insert Data", insertId)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO comments(email,comment) VALUES (?,?)"

	stmt, err := db.PrepareContext(ctx, script)
	if err != nil{
		panic(err)
	}

	defer stmt.Close()

	for i := 0; i <10 ; i++ {
		email := "jokowi" + strconv.Itoa(i) + "@gmail.comm"
		comment := "komentar ke " + strconv.Itoa(i)

		result, err := stmt.ExecContext(ctx,email,comment)
		if err != nil{
			panic(err)
		}

		id, err := result.LastInsertId()

		if err != nil{
			panic(err)
		}

		fmt.Println("id : ", id)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()

	if err != nil{
		panic(err)
	}
	script := "INSERT INTO comments(email,comment) VALUES (?,?)"

	for i := 0; i <10 ; i++ {
		email := "jokowi" + strconv.Itoa(i) + "@gmail.comm"
		comment := "komentar ke " + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx,script,email,comment)
		if err != nil{
			panic(err)
		}

		id, err := result.LastInsertId()

		if err != nil{
			panic(err)
		}

		fmt.Println("id : ", id)
	}
	err = tx.Commit()
	if err != nil{
		panic(err)
	}


}

func TestRepositoryTest(t *testing.T) {
	commentRepository := repository.NewCommentRepository(GetConnection())
	ctx := context.Background()
	comment := entity.Comment{
		Email: "aku@test.com",
		Comment: "test Repository",
	}

	result, err := commentRepository.Insert(ctx,comment)
	if err!=nil{
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := repository.NewCommentRepository(GetConnection())
	comment, err := commentRepository.FindById(context.Background(),4)
	if err!=nil{
		panic(err)
	}

	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	commentRepository := repository.NewCommentRepository(GetConnection())
	comments, err := commentRepository.FindAll(context.Background())
	if err!=nil{
		panic(err)
	}

	for _, comment:= range comments {
		fmt.Println(comment)
	}
}