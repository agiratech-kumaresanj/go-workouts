package main
import "fmt"

func main(){


 elements := map[string]map[string]string{

    "H": map[string]string{
      "name":"Hydrogen",
      "state":"gas",
    },
    
    }
/*	x := make(map[int]int)
	x[1] = 10
	delete(x,1)
	fmt.Println(x[1])
*/	
    fmt.Println(elements["H"])
}