package main

import ("fmt"
        "github.com/lib/pq"
        "database/sql"
        )

func main(){
	db, err := sql.Open("postgres", dataSourceName)

	if err := db.Ping(); err != nil {
    log.Fatal(err)
    	fmt.Println("not-connected")

    }else{
    	fmt.Println("connected")

    }

}