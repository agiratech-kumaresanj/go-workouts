package main
import "fmt"

var great int

func main(){
    //----------variable Declaration----------------
    var size1 int
    var size2 int
    var t bool = false
    //var x [10]int
     
    fmt.Println("Enter the 1st array size:")
    fmt.Scanf("%d", &size1)
    var array1 = make([]int,size1)
    
    fmt.Println("Enter the 2nd array size:")
    fmt.Scanf("%d", &size2)
    var array2 = make([]int,size2)
    //---------input values
        fmt.Println("Enter the Set of", size1, "values:")
    
	    for i :=0; i < size1; i++ {
	    
	     fmt.Scanf("%d",&array1[i])
	    
	    } 

	    fmt.Println("Enter the Set of", size2, "values:")
    
	    for i :=0; i < size2; i++ {
	    
	     fmt.Scanf("%d",&array2[i])
	    
	    }    
    //Check that values which one is greatest 
	        for i :=0; i < len(array1); i++ {
	        	for j :=0; j < len(array1); j++ {
                    if array1[i] == array2[j] {
                    	fmt.Println("Common values in Both array1 & array2:", array1[i])

                    } /*else {
                    	fmt.Println("UN-Common values in Both array1 & array2:", array1[i],)
                    }*/
	        	}
	           
	        }
	                    for i :=0; i < len(array1); i++ {
                                    t = false 
	        	           for j :=0; j < len(array2); j++ {
                               if array1[i] == array2[j] {
                    	            t = true
                               }
                            }
                               if(!t) {
                               fmt.Println("Un-Common values in array1:", array1[i])
                               }
                       }	    

	                    for i :=0; i < len(array2); i++ {
                                    t = false 
	        	           for j :=0; j < len(array1); j++ {
                               if array2[i] == array1[j] {
                               	    t = true
                               } 
                            }
                               if(!t){
                               fmt.Println("UN-Common values in array2:", array2[i])
                               }                        
                        }	    

	
}
	