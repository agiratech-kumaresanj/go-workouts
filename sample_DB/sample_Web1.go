package main

import (
	"fmt"
//	"io/ioutil"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

/*func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
*/
func viewHandler(w http.ResponseWriter, r *http.Request) {
	//title := r.URL.Path[len("/view/"):]
	//p, _ := loadPage(title)
//    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	var name, place string = "kumaresan","chennai"
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", name, place)
}

func main() {
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(":8080", nil)
}