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

var count int = 0
var le int = 0
var re int = 0
var child int = 0
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
    fmt.Println("Maximum_height of a tree is:", m_Depth(root))
    
    fmt.Println("Minimum_height of a tree is:", mi_Depth(root))
    
    fmt.Println("levels of a tree is:", level(root, n-1, 1))
    
    //fmt.Println(child,"child", count,"root",le,"left",re,"right")

}

func insert (t *tree, v int) *tree{
	count ++
	if(t == nil){
		t = &tree{v,nil,nil}
		child ++
	}else if(v<t.value){
		t.left = insert(t.left, v)
		 le ++
	}else {
		t.right = insert(t.right, v)
		 re ++
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

func m_Depth(d *tree) int{
	if (d == nil){
		return 0
	} else {
		l := m_Depth(d.left)
		r := m_Depth(d.right)
		if(l > r){
			return l+1
		}else{
			return r+1
		}
	}
}

func mi_Depth(d *tree) int{
	
	//var l, r bool
	if (d == nil){

		return 0

	}else if((d.left == nil) && (d.right == nil)) {

	    return 1

	}else if (d.left == nil) {

		return mi_Depth(d.right) + 1
		
	}else if (d.right == nil) {

		return mi_Depth(d.left) + 1
		
	}else{

		return Min(mi_Depth(d.right), mi_Depth(d.left)) + 1
	}

		
		
}


/*func level(l *tree) int{

  if (l == nil){
  	return lev
  }/*else if (l == l1 ){
  	return lev
  }else if ((l.left == nil) && (l.right == nil)) {
  	return -1
  	
  }else{
      
      l2 := level(l.left)
      r2 := level(l.right)

    if(l2 == -1){
  	   return r2
    }else{
       return l2  
    }
}

}*/

func level(l *tree, l1 int, lev int) int{

  if (l == nil){
  	return lev
  }else if (l.value == l1 ){
  	return lev+1
  }else{
  	return (level(l.left, l1, lev+2) | level(l.right, l1, lev+2))
  }

}

func Min(l int, r int) int{
	if(l > r){
		return r
	}else{
		return l
	}
}

