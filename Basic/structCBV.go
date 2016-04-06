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
	
	    fmt.Println("Enter emp_no:")
	    fmt.Scanf("%d", &d1.emp_no)
	    fmt.Println("Enter emp_name:")
	    fmt.Scanf("%s", &d1.emp_name)
	    fmt.Println(*d1)
        de = append(de, *d1)	
	     
	}
	fmt.Println(de)

	var t1 int
	fmt.Println("Enter you choice \n1.call by value:\n 2.call by reference")
	fmt.Scanf("%d",&t1)

	switch t1{
		case 1: fmt.Println("Call By value:")
             call_Byalue(de)
             fmt.Println(de) 
        case 2: fmt.Println("Call by reference:")
             call_ByReference(&de)
             fmt.Println(de)   

	}
    
    	
	
	 
}

func call_Byalue(call []company) {
	
	p := &company{}
	
	fmt.Println("Enter emp_no:")
	fmt.Scanf("%d", &p.emp_no)
	fmt.Println("Enter emp_name:")
	fmt.Scanf("%s", &p.emp_name)
    call = append(call, *p)
	fmt.Println("\ncallByValue the values:",call)

}


func call_ByReference(call *[]company) {
	
	p := &company{}
	fmt.Println("Enter emp_no:")
	fmt.Scanf("%d", &p.emp_no)
	fmt.Println("Enter emp_name:")
	fmt.Scanf("%s", &p.emp_name)
	*call = append(*call, *p)
	fmt.Println("\ncallByReference the values:",*call)

}


func main() {

	var emp int
	fmt.Println("Enter No/- of employee:")
	fmt.Scanf("%d",&emp)
	member(emp)
	
}

