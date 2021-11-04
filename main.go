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
	prev *Block
}
type Node struct {
	//Could store blocks in a map or in a linked list -
	name    int
	current Block
}

type log struct {
	nodes []Node
	current Block
}

var logToNode map[int]chan Block
var nodeToLog chan Block
var logger log

//Checks validity of block
func checkValid(){

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
func init(){
	numMiners := 4
	initBlock := Block{0,0,0, nil}
	logger.current = initBlock
	for i := 0; i < numMiners; i++{
		logger.nodes = append(logger.nodes, Node{i,initBlock})
	}
}

func main() {
	init()
}
