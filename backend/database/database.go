package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	UserName      string `json:"user_name"`
	UserID        int    `json:"user_id"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	MobileNummber string `json:"mobile"`
}

func DatabaseConnection() (*sql.DB, error) {
	//dbclient := "user=root dbname=restaurant sslmode=disable password=pass1234"
	dbclient := "root:pass1234@tcp(localhost:3306)/restaurant"
	db, err := sql.Open("mysql", dbclient)
	if err != nil {
		log.Fatal(err)

	}

	defer db.Close()

	db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The database connected successfully...")
	return db, nil
}
