package main

import (
    "fmt"
    "net/http"
    "./connection"
    //"html"
    "strconv"
    
)

var js string



func main() {
    //js = dbCon.DB_Connection()
    
   // http.HandleFunc("/", handler_html)
    connection.DB_con()
    http.Handle("/", http.FileServer(http.Dir("./public")))
    http.HandleFunc("/connection", handler)
    http.HandleFunc("/search", handler_search)
    
    http.ListenAndServe(":7007", nil)
    fmt.Println("------")

}

func handler(w http.ResponseWriter, r *http.Request){
   //fmt.Println(w)
    click := connection.DB_insert()
  
    fmt.Fprintf(w, click)
   
}

func handler_search(w http.ResponseWriter, r *http.Request){

   fmt.Fprintf(w, "Enter search ID:")
   fmt.Println(r.FormValue("name"))
   //value := r.URL.Path[1:]
   //fmt.Println(v.Get(r.FormValue("name")))
   
   //v := url.Values{}
   value := r.FormValue("name")
   value1,_ := strconv.Atoi(value)
   search := connection.DB_Search(value1)
   fmt.Fprintf(w, search)

}

