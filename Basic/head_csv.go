package main
import (
  "fmt"
  "os"
  "io"
  "encoding/csv"
  "bufio"
)

func main() {
  file_cs, err := os.Open("sample.csv")

//connection checking   
if err != nil {

    fmt.Println("Error:", err)
    return
  }
//file closing
  defer file_cs.Close()

//read the file content
  read := csv.NewReader(bufio.NewReader(file_cs)
  
for {

    record, err := read.Read()
// read.Send()
    if err == io.EOF {
      break
    } else if err != nil {
    
    fmt.Println("Error:", err)
    return
    }
   
// print the value
for i := range record {
      fmt.Print(" ", record[i])
}
    fmt.Println()
}


}
