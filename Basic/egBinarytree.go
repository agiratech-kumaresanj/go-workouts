package main

import "fmt"

type Node struct{
	value int
	left *Node
	right *Node	

}

func insert(root *Node,v int) *Node{
    if root == nil{
        root = &Node{v,nil,nil}
    } else if v<root.value{
        root.left = insert(root.left,v)
    } else{
        root.right = insert(root.right,v)
    }
    return root
}


func inTraverse(root *Node){
    if (root==nil){
        return
    }
    inTraverse(root.left)
    fmt.Printf("%d \t",root.value)
    inTraverse(root.right)
}


func main() {
    var treeRoot *Node
    var n int
    fmt.Println("Enter number of value:")
    fmt.Scanf("%d", &n)
    var a []int   
    a = make( []int, n) 
    for i := 0; i < n; i++{
    fmt.Scanf("%d", &a[i])
}
    
    fmt.Println("Array of integer: ")
    for i:=0;i<n;i++{
        fmt.Printf("%d ",a[i])
    }
    
    fmt.Println()
    for i:=0;i<n;i++{
       treeRoot = insert(treeRoot,a[i])
    }
    inTraverse(treeRoot)
    fmt.Println()
}