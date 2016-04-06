package main
import (
  "database/sql"
   _ "github.com/lib/pq"
  "fmt"
  // "log"
)



func InitDB(){
  db, dbErr = sql.Open("postgres", "user=postgres dbname=pqgotest password=root")

  if dbErr != nil {
    fmt.Println(dbErr)
  }
  fmt.Println("DB Connected")
}