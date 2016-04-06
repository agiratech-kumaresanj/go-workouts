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
    var number int = 0
    //VARIABLE
    var temp int = n
   //CONDITION 
    for temp != 0 {
      var rem int
  	  rem = temp%10
  	  number += rem*rem*rem
  	  temp = temp/10
    }
    //OUTPUT
    if n == number {
	fmt.Println("Armstrong Number", number)
    } 

 }
