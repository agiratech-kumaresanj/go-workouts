package main
import "fmt"

func zero(x *int) {
  *x = 0
}
func main() {
  x := 5
  zero(&x)
xPtr := new(int)
  one(xPtr)
  fmt.Println(*xPtr) // x is 1
  fmt.Println(x) // x is still 5
}

func one(xPtr *int) {
  *xPtr = 12	
}
/*func main() {
  xPtr := new(int)
  one(xPtr)
  fmt.Println(*xPtr) // x is 1
}*/