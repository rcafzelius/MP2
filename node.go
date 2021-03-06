package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
)

type Node struct {
	name      int
	current   Block
	logToNode chan Block
}

//
func (n *Node) mine(difficulty int) {
	//Store the current block for checking if the value has been changed later
	oldHead := n.current
	//Transaction in the new block states the miner has 10 coins
	transaction := fmt.Sprintf("%d=10", n.name)
	sum := sha256.New()
	sum.Write([]byte(transaction))
	//Create a hash of the transaction stored as a string
	hashTransaction := hex.EncodeToString(sum.Sum(nil))
	nonce := 1
	//Create a new block that points to the old block, stores the nonce, transaction and its hash, and an identifier
	newBlock := Block{oldHead.name + 1, nonce, hashTransaction, transaction, &oldHead}
L:
	//Check the validity of the Hash and if valid add the newBlock to the nodeToLog channel
	for {
		if checkValid(newBlock, difficulty) {
			nodeToLog <- newBlock
			break L
			//Breaks the mining loop if n.current ever changes.
		} else if oldHead != n.current {
			break L
		}
		nonce = rand.Intn(MaxInt)
		newBlock.nonce = nonce
		//Simulate faulty behavior by sending a probabilistically incorrect block hash into the nodeToLog channel
		if nonce <= MaxInt/1000000000 {
			nodeToLog <- newBlock
		}
	}
}

//Check the logToNode channel and update the node current to newBlock
func (n *Node) listen() {
L:
	for {
		select {
		case newBlock, ok := <-n.logToNode:
			if ok {
				if n.current != newBlock {
					n.current = newBlock
				}
				break L
			}
		default:
		}
	}
}
