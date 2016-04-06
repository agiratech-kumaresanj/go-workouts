package main
import "fmt"

type name interface{
	class() int
}

func main(){
	t *name
	fmt.Println("interface")
	fmt.Println(t.class())
}