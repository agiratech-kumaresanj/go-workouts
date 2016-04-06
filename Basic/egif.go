package main
import "fmt"

var x [5]int
var count int =0
var total int =0
func main() {
   
  fmt.Println("Enter the 5 values:")
 
  for i :=0; i < 5; i++{
  	fmt.Println(" position of value:",i)
  	fmt.Scanf("%d",&x[i])
  }
   fmt.Println("Total:")
   totalvalue()
   fmt.Println("Average value:")
   average()
  
}

func totalvalue(){
	for i := 0; i < 5; i++ {
    total += x[i]
  }
  fmt.Println(total)
}

func average(){
	for i := 0; i < 5; i++ {
    total += x[i]
    count++
  }
  fmt.Println(total/count)
}