package main
import "fmt"

type prin interface{
	prin1() 
}

type tree struct{
    value int
    left *tree
    right *tree

}

func main(){
	
	var n int
	fmt.Println("enter value:")
	
	fmt.Scanf("%d", &n)
	var a []int = make([]int,n)
	
	for i:=0; i < n; i++{
		fmt.Scanf("%d", &a[i])
	}
	var root *tree
	fmt.Println("Array of Integers value:")
	for i:=0; i < n; i++{
		fmt.Printf("%d \t", a[i])
	}
    
    fmt.Println()
    for i :=0; i < n; i++{
        root = insert(root, a[i])
    }
    
    side(root)
    fmt.Println()
    
    fmt.Println("Are you want to insert new value:press: 'y'")
    var yes string
    fmt.Scanf("%s", &yes)
    
    if(yes == "y"){
     	var new int
     	fmt.Println("Enter the value:")
     	fmt.Scanf("%d", &new)
     	root = insert(root,new)
     }

    side(root)
    fmt.Println()


}

func insert (t *tree, v int) *tree{
	
	if(t == nil){
		t = &tree{v,nil,nil}
	}else if(v<t.value){
		t.left = insert(t.left, v)
		
	}else {
		t.right = insert(t.right, v)
	}
	
	
	return t

}

func side (t *tree){
	
	if (t == nil){
		return
	}
	side(t.left)
	t.prin1()
	side(t.right)
	
}

func (t *tree) prin1()  {
	
	fmt.Printf("%d \t", t.value)
	
}