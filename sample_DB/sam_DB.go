package main
import (
  "database/sql"
   _ "github.com/lib/pq"
  "fmt"
  // "log"
)

//var db *sql.DB
//var dbErr error

//const timeLayout = "Jan 2, 2006"

//const tableName = "testcsv"


func main(){
InitDB()	
}
func InitDB(){
	db, dbErr := sql.Open("postgres", "user=postgres password='root' dbname='sample' sslmode=disable")

  if dbErr != nil {
    fmt.Println(dbErr)
  }
  fmt.Println("DB Connected")
  fmt.Println(db)
}