package main
import "fmt"

//var n int
var number int = 0
var rem int
func main(){
  
    //----------variable Declaration----------------
    var size int
    fmt.Println("Enter the limited Number:")
    fmt.Scanf("%d", &size)

    for i :=0; i <= size; i++ {

    	arm_Strong(i)

    } 

}    


 func arm_Strong(n int){
	
    /*var n int
    fmt.Println("Enter the Number:")
    fmt.Scanf("%d", &n)
    */
    var temp int = n

   /* var size1 int
    var size2 int
 
    //var x [10]int
     
    fmt.Println("Enter the 1st array size:")
    fmt.Scanf("%d", &size1)
    var array1 = make([]int,size1)
*/

      
	    for temp != 0 {

	  	  rem = temp%10
	  	  number += rem*rem*rem
	  	  temp = temp/10

	    }
	    //    fmt.Println("Result:")

	        if n == number {
	    	fmt.Println("Armstrong Number", number)
	        } 
 	
 }
