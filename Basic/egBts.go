package main
import "fmt"

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

}

func insert (t *tree, v int) *tree{
	
	if(t == nil){
		t = &tree{v,nil,nil}
	}
	return t

}

func side (t *tree){
	
	if (t == nil){
		return
	}
	
	side(t.left)
	fmt.Printf("%d \t", t.value)
	side(t.right)
	
}

func search(s *tree, v){
	
	if(v < t.value){
		t.left = insert(t.left, v)
	}else if{
		t.right = insert(t.right, v)
	}else {
		search(t)
	}
}
