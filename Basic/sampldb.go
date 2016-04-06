package main

import (
  "encoding/csv"
  "fmt"
  "io"
  "os"
)

func main() {
  file, err := os.Open("sample.csv")
  var x int =0

 if err != nil {
    // err is printable
    fmt.Println("Error:", err)
    return
  }
  // automatically call Close() at the end of current method
  defer file.Close()
  //
  reader := csv.NewReader(file)
  // options are available at:
  // http://golang.org/src/pkg/encoding/csv/reader.go?s=3213:3671#L94
  reader.Comma = ';'
  lineCount := 0
  for {
    // read just one record, but we could ReadAll() as well
    record, err := reader.Read()
    // end-of-file is fitted into err
    if err == io.EOF {
      break
   } else if err != nil {
      fmt.Println("Error:", err)
      return
    }

    //skip header line in csv
    if x == 0 {
      fmt.Println("Error")
      x++
      continue
    }

    // record is an array of string so is directly printable
    // fmt.Println("Record", lineCount, "is", record, "and has", len(record), "fields")
    // and we can iterate on top of that
    for i := 0; i < len(record); i++ {
      fmt.Println(" ", record)
    }
    fmt.Println()
    lineCount += 1
  }
} 