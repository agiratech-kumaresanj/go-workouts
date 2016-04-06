package main
import "fmt"

func main(){

	var size1 int
	fmt.Println("Enter a size1")
	fmt.Scanf("%d", &size1)
	var array1 = make([]int,size1)
    
     x := make(map[int]int)
     //keys := make([]int, 0, len(x))
	
	for i := 0; i < len(array1); i++{
		//fmt.Println("value:")
		fmt.Scanf("%d", &array1[i])

	}

	fmt.Println("..........condition...........\n")
    var count int = 1
    fmt.Println("OUT:",array1[:])

    for i := 0; i < len(array1); i++ {
    	
    	for j := i+1; j < len(array1); j++ {

    		if array1[i] == array1[j] {   
    		   array1[j] = '*'         
                count++
            }
    	}
        
        if count > 1 && array1[i] != '*' {
    	
    	//fmt.Println("        ======        ", count)
   // 	fmt.Println("        ======        ")
    	
    	x[array1[i]] = count
    	fmt.Println(array1[i]," ----> ", x[array1[i]])
    	
        
        }
        count = 1
    }

    /*fmt.Println("OUT")
    for i := 0; i < len(x); i++ {
    	fmt.Println(i,"  ", x[i])
    }*/
}