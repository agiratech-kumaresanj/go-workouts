package main

import (
    "fmt"
	"database/sql"
    "log"

    )

func main() {
	
	db_Connection()

/*	fmt.Println("enter: 1.yes or 2.no")
	fmt.Scanf("%s", &con)
	db_Connection()
	if(con == "yes"){
    }else{
    	main()
    }*/
}

func db_Connection(){

	db, err := sql.Open("postgres", "user=pqgotest dbname=pqgotest sslmode=verify-full")
	
	if err != nil {
		log.Fatal(err)
		
		//fmt.Println("there are errors:%v", err)

		fmt.Println("there are errors")

	}else {
		fmt.Println("no errors")
	}
	defer db.Close()
 fmt.Println(db)


}
