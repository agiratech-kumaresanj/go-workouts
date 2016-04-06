package main
import "fmt"

type Product struct {
	name string
	code int
	price float64
}

func main() {
	var num int
	fmt.Println("How many elements you want to enter :")
	fmt.Scanf("%d",&num)
	takingInput(num)
}

func takingInput(num int) {
	p := &Product{}
	a := []Product{}

	for i:=0;i<num;i++ {
		fmt.Println("Enter the Product Details")
		fmt.Printf("ProductName :")
		fmt.Scanf("%s",&p.name)
		fmt.Printf("\nProductCode : ")
		fmt.Scanf("%d",&p.code)
		fmt.Printf("\nProductPrice : ")
		fmt.Scanf("%f",&p.price)
		a = append(a, *p)
	}
	fmt.Println("\nBefore calling callByVal the values are :",a)
	callByValue(a)
	fmt.Println("After calling the values are :",a)
	fmt.Println("================================")
	fmt.Println("Before calling callByRef the values are",a)
	callByReference(&a)
	fmt.Println("After calling the values are :",a)
}

func callByValue(a []Product) {
		p := &Product{}
		fmt.Println("Please enter one more Product Details")
		
		fmt.Printf("ProductName :")
		fmt.Scanf("%s",&p.name)
		fmt.Printf("\nProductCode : ")
		fmt.Scanf("%d",&p.code)
		fmt.Printf("\nProductPrice : ")
		fmt.Scanf("%f",&p.price)
		a = append(a, *p)
		fmt.Println("\nIn callByValue the values are :",a)
}

func callByReference(a *[]Product) {
		p := &Product{}
		fmt.Println("Please enter one more Product Details")
		
		fmt.Printf("ProductName :")
		fmt.Scanf("%s",&p.name)
		fmt.Printf("\nProductCode : ")
		fmt.Scanf("%d",&p.code)
		fmt.Printf("\nProductPrice : ")
 		fmt.Scanf("%f",&p.price)
 		*a = append(*a, *p)
 		fmt.Println("\nIn callByValue the values are :",*a)
 }