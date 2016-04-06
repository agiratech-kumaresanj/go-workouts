package main

import "fmt"
//import "math"

type access interface{
	swipe_in() string
	swipe_out() string
}

type team1 struct{
	lang string
	lang1 string
}
type team2 struct{
	lang3 string
    lang4 string
}
func (t team1) swipe_in() (string, string){
	return * t.lang * t.lang1
}

func (t2 *team2) swipe_out() (string, string) {
	return t2.lang3, t2.lang4
}

func company(a access) {
	fmt.Println(a)
	fmt.Println(a.swipe_in())
	fmt.Println(a.swipe_out())
}
func main(){
    y := team1 {lang : "Golang", lang1 : "Ruby" }
    n := team2 {lang3 : "Php", lang4 : "javaScript"}

    company(y)
    company(n)
}