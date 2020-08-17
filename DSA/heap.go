package main

import (
	"fmt"
)

func heapify(x []int){

	for i:=len(x)-1;i>0;i--{
		parent := int((i-1)/2)
		c1 := parent*2 + 1
		c2 := parent*2 + 2
		
		for {
			
			toswap := x[parent]
			if(x[c1] > toswap){
				toswap = x[c1]
			}
			if( c2 < len(x) && x[c2] > toswap){
				toswap = x[c2]			
			}
			
			if( toswap == x[c1]){
				x[c1],x[parent] = x[parent],x[c1]
				parent = c1								
			} else if (c2 < len(x) && toswap == x[c2]){
				x[c2],x[parent] = x[parent],x[c2]
				parent = c2
			} else {
				break
			}	
			
			if( parent*2 + 1 <len(x) ) {
				c1 = parent*2 + 1
				c2 = parent*2 + 2
			} else {
				break
			}	
			
			_ = toswap
			
		}

	}
}

func main() {
	var arr []int = []int{11,90,32,50,-12,9,-45,420,-69}
	heapify(arr)
	for i:=0;i<len(arr);i++ {
		fmt.Println(arr[i])
	}
}
