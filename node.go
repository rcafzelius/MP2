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

func (n *Node) mine(difficulty int) {
	oldHead := n.current
	transaction := fmt.Sprintf("%d=10", n.name)
	sum := sha256.New()
	sum.Write([]byte(transaction))
	hashTransaction := hex.EncodeToString(sum.Sum(nil))
	nonce := 1
	newBlock := Block{oldHead.name + 1, nonce, hashTransaction, transaction, &oldHead}
L:
	for {
		if checkValid(newBlock, difficulty) {
			nodeToLog <- newBlock
			break L
		} else if oldHead != n.current {
			break L
		}
		nonce = rand.Intn(MaxInt)
		newBlock.nonce = nonce
	}
	println("found")
}

func (n *Node) listen() {
	//while loop listens to channel
	//update current block in struct
L:
	for {
		select {
		case newBlock, ok := <-n.logToNode:
			if ok {
				n.current = newBlock
				break L
			}
		default:
		}
	}
	println("heard")
}
