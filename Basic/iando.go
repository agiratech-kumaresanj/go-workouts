package main

import "fmt"

func main() {
  fmt.Println("Enter a Details: ")
  
  
  fmt.Print("Enter a Name: ")
  var name string
  fmt.Scanf("%s", &name)

  fmt.Print("Enter a Age: ")
  var age int
  fmt.Scanf("%d", &age)

  fmt.Print("Enter a Place: ")
  var place string
  fmt.Scanf("%s", &place)
  

   
   
   fmt.Printf("Hi %s Your age is %d and You are from %s \n", name, age, place)

}
