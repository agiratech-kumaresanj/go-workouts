package main

import "fmt"
    


func main() {
    // To create a map as input
    m := make(map[int]string)
    m[1] = "a"
    m[2] = "c"
    m[0] = "b"
    b := make(map[int]string)
    
    var c int
    fmt.Println("choose your choice:")
    fmt.Scanf("%d", &c)
    switch c{
       case 1:
              fmt.Println("You need to insert a value:")
             b = push(m)
                // push(m)
            /* for i := 0; i < len(b); i++{
                if ()
             }*/
              fmt.Println("final value:", b)
       case 2:
             b = pop(m)
        default :
            fmt.Println("Enter right choice:")
            main()
    }  

    /*for _, k := range keys {
        fmt.Println("Key:", k, "Value:", m[k])*/
    
}

func push(a map[int]string) map[int]string {
    fmt.Println(a)
    var name string
    var keys int
    fmt.Println("Need to insert a key & value:")
    fmt.Scanf("%d", &keys)
    fmt.Scanf("%v", &name)
    a[keys] = name
    fmt.Println(a)
   return a
}

func pop(a map[int]string) map[int]string {

    var keys int
    fmt.Println(a)
    fmt.Println("Enter delete value:")
    fmt.Scanf("%d", &keys)
    delete(a,keys)
    fmt.Println(a)
    return a

}