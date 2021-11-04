package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type Block struct {
	name, hash, solution int
}

type Node struct {
	//Could store blocks in a map or in a linked list --
	blocks map[int]Block
}

var logToNode map[int]chan Node
var nodeToLog chan int


//Checks validity of block
func logger(){

}

//Updates nodes when a new valid block is added
func updateNodes() {

}

//checkStatus after each iteration in solver check for an update of chain
func checkStatus() {

}

//Take in Node and channel to send a solution
func solver(){

}

//Run Go Routines to have nodes run a mining function "solver" at the same time
func mine() {
	for
}

func main() {

}
