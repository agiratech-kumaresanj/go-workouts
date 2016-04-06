package main

import (
  
  "net/http"
)

func main() {
  http.Handle("/", http.FileServer(http.Dir("./public")))
 
  http.ListenAndServe(":8000", nil)
}
// func serveTemplate(w http.ResponseWriter, r *http.Request) {
//   lp := path.Join("templates", "layout.html")
//   fp := path.Join("templates", r.URL.Path)

//   tmpl, _ := template.ParseFiles(lp, fp)
//   tmpl.ExecuteTemplate(w, "layout", nil)
// }