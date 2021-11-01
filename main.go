package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)




type Block struct {
	name, puzzle int

}

//
type Node struct {
	//Could store blocks in a map or in a linked list --
	blocks map[int]Block
	solution map[int]int

	sendBlock map[int]chan Block
}



//Checks validity of block
func logger(){

}

//Updates nodes when a new valid block is added
func updateNodes() {

}

func main() {

}
