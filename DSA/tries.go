package main

import (
	"fmt"
)

var LIM = 26
type trie struct{
	isTerminal bool
	nextWord[] trie
	isInit bool
}


func main() {

	var myTrie trie ;
	var searchWords []string = []string{"batman","superman","wonderwoman","starfire"}
	
	myTrie.isInit = true
	myTrie.nextWord = make([]trie,LIM,LIM)
		
	
	for i:= range searchWords {
		ref:=&myTrie
		for j:= range searchWords[i]{
			location := int(searchWords[i][j]) - int('a')
						
			if (!ref.nextWord[location].isInit){
				ref = &ref.nextWord[location]
				ref.isInit = true
				ref.nextWord = make([]trie,LIM,LIM)
			} else{
				ref = &ref.nextWord[location]				
			}
		}		
		ref.isTerminal = true
	}

	
	//Search part
 	var text = "what if batman and starfire were literal in terms of meaning"
	
	for i:= range text {
		searchTrie(i,text,myTrie);	
	}
	

}

func searchTrie(startLoc int, text string, myTrie trie){
	i:= startLoc
	ref:= myTrie
	var loc = int(text[i]) - int('a')
	
	if (loc == int(' ') - int('a')){
		return
	}
	
	for ref.nextWord[loc].isInit && i<len(text){
		if(ref.isTerminal){
			fmt.Println(text[startLoc:i])
			break
		}else{
			ref = ref.nextWord[loc]
			i++
			loc= int(text[i]) - int('a')
			if (loc == int(' ') - int('a')){
				break
			}	
		}		
	}
	if(ref.isTerminal){
		fmt.Println(text[startLoc:i])
		return
	}	
}