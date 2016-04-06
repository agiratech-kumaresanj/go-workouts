package main

import (
	  "database/sql"
	_ "github.com/go-sql-driver/mysql"
      "log"
      "fmt"
)

func main() {
	db, err := sql.Open("mysql",
		"user:password@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		log.Fatal(err)
		fmt.Println("not connected")
	}else{
		fmt.Println("not connected")
		
	}
	defer db.Close()
}