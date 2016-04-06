package main

import (
    "fmt"
    "net/http"
    "./dbCon"
    //"html"
)

var js string

/*
func handler_html(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, dbCon.DB_Connection)
    fmt.Fprintf(w, "Hello astaxie!")
    fmt.Println(w)
// write data to response
}*/



func main() {
    js = dbCon.DB_Connection()

    http.Handle("/", http.FileServer(http.Dir("./public")))
    http.HandleFunc("/dbCon", handler)

    http.ListenAndServe(":8080", nil)
    fmt.Println("------")

}

func handler(w http.ResponseWriter, r *http.Request){
   fmt.Println(w)
    w.Write([]byte(js))
}


/*func main() {
    
    //	var s []Declare
    fmt.Println(http.Dir("html"))
    fs := http.FileServer(http.Dir("/html"))
    http.Handle("/",fs)
    //http.HandleFunc("/html", handler_html)
    http.HandleFunc("/deCon", handler)
    http.ListenAndServe(":8080", nil)
     dbCon.DB_Connection()

    //dbCon.Convert_Json(s)
    //s = dbCon.DB_Connection()
    //fmt.Println(s)
    //dbCon.Convert_Json(d_store[]Declare)
}
*/