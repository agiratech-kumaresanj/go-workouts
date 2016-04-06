package main
import "fmt"

func callBy_Value(add1 *int) {
	fmt.Println("====Inside of process callBy_Value====")
	fmt.Println("1.value_add:", *add1)
	*add1 +=100
	fmt.Println("2.value_add:", *add1)
}

func main() {

	fmt.Println("=========call by reference=========")
	var add int 
	fmt.Println("Enter your value")
    fmt.Scanf("%d", &add)
	fmt.Println("1.add: ", add)
	fmt.Println("Before variable value:")
	callBy_Value(&add)
	fmt.Println("after variable value:")
	fmt.Println("3.add:", add)
}  