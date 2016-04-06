package main

import "fmt"

//Linked list structure
type Single_inter struct{
	
	value interface{}
	next *Single_inter

}

type Single struct{

	top *Single_inter
    size int
}
//Main funcation

func main(){
	s := new(Single)
	var st [3]string
	var cases1 int = 1
	//fmt.Scan(cases1) 
	for cases1 != 0{
	fmt.Println("Enter switch case choice 1 and 2 :")
    var cases int
	fmt.Scanf("%d",&cases)
	switch(cases){
	case 1:
		fmt.Println("push: ")
		for i :=0; i < 3; i++{
		fmt.Scanf("%s", &st[i])
	    s.push(st)
        }
    case 2:
    	for i :=0; i < 3; i++{
		//fmt.Scanf("%s", &st[i])
    		fmt.Println(st[i])
	    s.pop(st)
        }
     default :
        cases1 = 0   
	}
	/*for s.size > 0 {
		fmt.Printf("%s ", s.pop().(string))
	}
	    fmt.Println()*/ 
	}

//	fmt.Println(s.push(st))
}
//push function

func (s *Single) push(value interface{}) {
	
	s.top = &Single_inter{value, s.top}
	s.size ++
	//fmt.Println(s.top)
	return
	/*var t string
	return t := s.top*/ 
}

func (s *Single) pop(value interface{}) {
	
	if s.size > 0{
     value, s.top = s.top.value, s.top.next
     fmt.Println("", s.top, value)
     s.size --
     return
	}
		return  

}