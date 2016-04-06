package main
import (
  "fmt"
  "os"
  "io"
  "encoding/csv"
  "bufio"
  "log"
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
  read := csv.NewReader(bufio.NewReader(file_cs))
  
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
// write
    record,_ := read.Read()

write := csv.NewWriter(os.Stdout)
    
      for _, record := range record {
        if err := write.Write(); err != nil {
          log.Fatalln("error writing record to csv:", err)

        }else{
          fmt.Println(record)
        }
      }
    
     // Write any buffered data to the underlying writer (standard output).
      write.Flush()
    
      if err := write.Error(); err != nil {
        log.Fatal(err)
      }


}


/*
  func ExampleWriter() {
    records := [][]string{
        {"first_name", "last_name", "username"},
        {"Rob", "Pike", "rob"},
        {"Ken", "Thompson", "ken"},
        {"Robert", "Griesemer", "gri"},
      }
    
      w := csv.NewWriter(os.Stdout)
    
      for _, record := range records {
        if err := w.Write(record); err != nil {
          log.Fatalln("error writing record to csv:", err)
        }
      }
    
     // Write any buffered data to the underlying writer (standard output).
      w.Flush()
    
      if err := w.Error(); err != nil {
        log.Fatal(err)
      }
      // Output:
      // first_name,last_name,username
      // Rob,Pike,rob
      // Ken,Thompson,ken
      // Robert,Griesemer,gri
    }


  func ExampleReader() {
    in := `first_name,last_name,username
    "Rob","Pike",rob
    Ken,Thompson,ken
   "Robert","Griesemer","gri"
    `
      r := csv.NewReader(strings.NewReader(in))
    
      for {
        record, err := r.Read()
        if err == io.EOF {
          break
        }
        if err != nil {
          log.Fatal(err)
        }
    
        fmt.Println(record)
      }
     // Output:
      // [first_name last_name username]
      // [Rob Pike rob]
      // [Ken Thompson ken]
      // [Robert Griesemer gri]
    }*/
    