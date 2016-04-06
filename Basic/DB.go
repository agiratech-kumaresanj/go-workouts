package main
 
import (
  //"os"
  "log"
  "time"
  "fmt"
  "database/sql"
 // _ "github.com/lib/pq"
)
 
func main() {
  // create the statement string
  var sStmt string = "insert into test (gopher_id, created) values ($1, $2)"
 
  // lazily open db (doesn't truly open until first request)
  db, err := sql.Open("postgres","host=localhost dbname=testdb sslmode=disable")
  if err != nil {
    log.Fatal(err)
  }
 
  stmt, err := db.Prepare(sStmt)
  if err != nil {
    log.Fatal(err)
  }
 
  fmt.Printf("StartTime: %v\n", time.Now())
 
  res, err := stmt.Exec(1, time.Now())
  if err != nil || res == nil {
    log.Fatal(err)
  }
 
  // close statement
  stmt.Close()
 
  // close db
  db.Close()
 
  fmt.Printf("StopTime: %v\n", time.Now())
}