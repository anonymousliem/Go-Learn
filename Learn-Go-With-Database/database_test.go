package Learn_Go_With_Database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"time"
)


func TestEmpty(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(0.tcp.ngrok.io:16385)/gobase")
	if err != nil {
		panic(err)
	}
	fmt.Println("halo", db)
}

func GetConnection() *sql.DB{
	db, err := sql.Open("mysql", "root:@tcp(0.tcp.ngrok.io:16385)/gobase?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5*time.Minute)
	db.SetConnMaxLifetime(1*time.Hour)

	return db
}
