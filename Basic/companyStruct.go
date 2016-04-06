package main
import "fmt"


type company struct{
	emp_no int
	emp_name string
}

func member(d int) {

	d1 := &company{} 
    de := []company {}
	 
	for i := 0; i < d; i++ {
	
	    fmt.Println("enter emp_no:")
	    fmt.Scanf("%d", &d1.emp_no)
	    fmt.Println("enter emp_name:")
	    fmt.Scanf("%s", &d1.emp_name)
	    fmt.Println(*d1)
        de = append(de, *d1)	
	     
	}
	
	fmt.Println(de)
	 
}


func main() {

	var t int
	fmt.Println("enter:")
	fmt.Scanf("%d",&t)
	member(t)
	switch t{
		case 1: fmt.Println("call By value:")
             member(t)
        case 2: fmt.Println("call by reference:")
             member(t)   
	}

}

