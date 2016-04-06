package main
import "fmt"
var temp int
func main(){

	x := []int{
  48,96,86,68,
  57,82,63,70,
  37,34,83,27,
  19,97,9,17,12,
}
    temp = x[0]
	    for v :=0; v < len(x); v++{
	//fmt.Println("",v)
           if temp < x[v]{
           	temp = x[v]
           }

			
		   }
        
    
    fmt.Println("Small:",temp)
}
	/*var a [10]int
	var n int
	fmt.Println("Enter N")
     fmt.Scanf("%d",&n)
	for i := 0; i < n; i++{
	
	fmt.Scanf("%d:",&a[i])
}*/