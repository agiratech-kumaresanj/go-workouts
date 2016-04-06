package main
import "fmt"

func main() {

  var i int
  fmt.Println("Enter ur choice")
  fmt.Println("1.append ")
  fmt.Println("2.copy")
  fmt.Scanf("%d",&i)
  
  if i == 1{ 
 
  fmt.Println("1.append ")
  slice1 := []int{1,2,3}
  
  fmt.Println("before:")
  fmt.Println(slice1)
  
  fmt.Println("after:")
  slice2 := append(slice1, 4, 5)
  fmt.Println(slice2)
  
  } else if i == 2 {
  
  fmt.Println("2.copy")
  slice3 := []int{1,2,3}
  slice4 := make([]int, 2)
  
  fmt.Println("!. before copy")
  fmt.Println(slice4, slice3)
  
  copy(slice4, slice3)
  fmt.Println("!!. after copy")
  fmt.Println(slice4, slice3)
  
  } else{
  	main()
  }
  

}

