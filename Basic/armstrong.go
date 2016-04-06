package main
import "fmt"

//var n int
func main(){
    //var a int
    //----------variable Declaration----------------
    var size int
    fmt.Println("Enter the limited Number:")
    fmt.Scanf("%d", &size)
    for i := 0; i <= size; i++ {
    	arm_Strong(i)
    } 
}    


 func arm_Strong(n int){
    
    //VARIABLE
    var temp int = n
    var re int 
   //CONDITION 
    re = recur(temp)
    if n == re {
	fmt.Println("Armstrong Number", n)
    } 

 }
//recursion
 func recur (temp int) int {
    var rem int
    var number int = 0
    
    rem = temp%10
    number += rem*rem*rem
    temp = temp/10
    
    if temp != 0 {
        recur(temp)
    } 
    return number
 }