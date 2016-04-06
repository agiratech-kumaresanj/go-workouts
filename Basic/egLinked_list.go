package main
 
 import (
 	"container/list"
 	"fmt"
 )
 
 func main() {
 	// Create a new list and put some numbers in it.
 	l := list.New()
 	e4 := l.PushBack(4)
 	e1 := l.PushFront(1)
 	l.InsertBefore(3, e4)
 	l.InsertAfter(2, e1)
 	var v,p int
 
 	fmt.Println("list of values :")
 	for e := l.Front(); e != nil; e = e.Next() {
 		fmt.Println(e.Value)
 	}
 	fmt.Println("number of elements in the list(length) :",l.Len())
 	
 	fmt.Println("enter which element you want to remove in the list :")
 	fmt.Scanf("%d",&v)
 			for e := l.Front(); e !=nil; e =e.Next() {
 				if(e.Value == v) {
 					l.Remove(e)
 				}
 			}
 	fmt.Println("after removing the list of values are :")
 			for e := l.Front(); e != nil; e = e.Next() {
 				fmt.Println(e.Value)
 			}
 	fmt.Println("which value after you want to insert :")
 		fmt.Scanf("%d",&p)
 	fmt.Println("Please enter the value :")
 		fmt.Scanf("%d",&v)
 			for e := l.Front(); e !=nil; e =e.Next() {
 				if(e.Value == p) {
 					l.InsertAfter(v,e)
 				}
 			}
 	fmt.Println("after inserting the list of values are :")
 			for e := l.Front(); e != nil; e = e.Next() {
 				fmt.Println(e.Value)
 			}
 
 