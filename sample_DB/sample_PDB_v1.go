package psqlConn

import (
  "database/sql"
   _ "github.com/lib/pq"
  "fmt"
  // "log"
)

var db *sql.DB
var dbErr error

const timeLayout = "Jan 2, 2006"

const tableName = "testcsv"



func InitDB(){
  db, dbErr = sql.Open("postgres", "user=postgres dbname=sample password=root")

  if dbErr != nil {
    fmt.Println(dbErr)
  }
  fmt.Println("DB Connected")
}

func checkLast(index,len int) string{
  var operator string
  if operator = ")"; index < len { operator = ","}
  return operator
}

func CreateTable(headers []string){
  check_table := fmt.Sprintf("create table IF NOT EXISTS %v",tableName)
  set_primary_key := "(ID integer PRIMARY KEY,"
  set_variables := ""
  // stringVal :=
  length  := len(headers)
  for i,header := range headers {
    // fmt.Println(header)
    value := fmt.Sprintf("%v TEXT%v",header,checkLast(i+1,length))
    set_variables += value
  }
  finalStr := check_table+set_primary_key+set_variables
  _, err := db.Exec(finalStr)
  if err != nil {
    fmt.Println(err)
  }
}

func InsertRec(rec string) {
  insertQuery := fmt.Sprintf("INSERT INTO %v values(%v)",tableName,rec)
  fmt.Println(insertQuery)
  _, err := db.Exec(insertQuery)
  if err != nil {
    fmt.Println(err)
  }
}