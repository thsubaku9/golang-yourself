package main

import( 
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int){
	if(t.Left != nil){
		Walk(t.Left,ch)
	}
	ch <- t.Value
	if(t.Right != nil){
		Walk(t.Right,ch)
	}	
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree,size int) bool{
	c1 := make(chan int)
	c2 := make(chan int)
	go Walk(t1,c1)
	go Walk(t2,c2)
	for i:=0;i<size;i++{
		v1:=<-c1
		v2:=<-c2
		if(v1!=v2){
			return false
		}
	}
	return true
}

func main() {

	fmt.Println(Same(tree.New(2),tree.New(2),10))
}
