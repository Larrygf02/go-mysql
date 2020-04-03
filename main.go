package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Go mysql Tutorial")

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/gotest")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	fmt.Println("Succesfully connected to mysql")

	insert, err := db.Query("Insert into users(name) values ('Josue')")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

	fmt.Println("Succesfully insert")
}
