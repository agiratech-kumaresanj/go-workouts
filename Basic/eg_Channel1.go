package main

import ("fmt"
        "time"
)

 func main() {
 var ch chan string = make(chan string)
 var ch1 chan int = make(chan int)

   go student_in_Detail(ch, ch1)
   go student_ou_Detail(ch, ch1)
var input float32
  fmt.Scanln(&input)
}

func student_in_Detail(ch chan string, ch1 chan int){

   var name string //= "ken"
   var ro_no int //= 10
   fmt.Scanf("%s", &name)
   fmt.Scanf("%d", &ro_no)
   for i := 0; ; i++{
   ch <- name//"ken"//name
   ch1 <- ro_no

	fmt.Println("geted value:", name)
	fmt.Println("geted value:", ro_no)

   time.Sleep(10000)
}
	//ch <- ro_no 
}

func student_ou_Detail(ch chan string, ch1 chan int) {
	pr := <- ch
	pr1 := <- ch1
    //fmt.Println(pr)
	fmt.Println("geted value:", pr)
	fmt.Println("geted value:", pr1)

}