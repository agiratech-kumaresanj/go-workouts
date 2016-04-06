package main

import "fmt"
var key string
var count int = 0
func main() {

  elements := map[string]map[string]string{

    "k": map[string]string{
      "name":"kumaresan",
      "state":"TN",
    },

    "R": map[string]string{
      "name":"Reddy",
      "state":"AP",
    },

    "A": map[string]string{
      "name":"Arun",
      "state":"AP&TN",
    },

    "S": map[string]string{
      "name":"Sathis",
      "state":"TN",
    },

    "Ke":  map[string]string{
      "name":"Ken",
      "state":"TN",
    },

  }
  fmt.Println("Enter the key value:")
  fmt.Scanf("%v", &key)
  if el, t := elements[key]; t {
    fmt.Println(el["name"], el["state"])
    count++
  
  }
  for count < 5 {
     main()
    
    } 
    
}