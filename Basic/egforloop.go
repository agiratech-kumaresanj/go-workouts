package main

import "fmt"

var num1, num2 int

func main() {
    
    fmt.Println("Enter the two Numbers:")
    fmt.Println("num1:")
    fmt.Scanf("%d", &num1)
    fmt.Println("num2:")
    fmt.Scanf("%d", &num2)
    fmt.Println("Enter the Number B/W 0-5:")  
    
    var i int
    fmt.Scanf("%d",&i)
    
    switch i {
        case 0: fmt.Println("Zero")
        case 1: fmt.Println("ADD")
            add()
        case 2: fmt.Println("SUB")
            sub()
        case 3: fmt.Println("MUL")
            mul()
        case 4: fmt.Println("DIV")
            div()
        case 5: fmt.Println("Five")
            main()
        default: fmt.Println("Unknown Number")
    }

    if i < 6 {
        main()	
    }

}

    func add() {
    	
    	fmt.Println("The addition of two Number:")
    	num1 +=num2
    	fmt.Println(num1)

    }

    func sub(){
    	
    	fmt.Println("The Substraction of two Number:")
    	num1 -=num2
    	fmt.Println(num1)

    }

    func mul(){
    	
    	fmt.Println("The Multification of two Number:")
    	num1 *=num2
    	fmt.Println(num1)

    }

    func div(){
    	
    	fmt.Println("The Divition of two Number:")
    	num1 /=num2
    	fmt.Println(num1)

    }