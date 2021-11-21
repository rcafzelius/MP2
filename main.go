package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

const MaxInt = int(^uint(0) >> 1)

type Block struct {
	name, nonce int
	hash        string
	//add single transaction as map: {source: location coin came from, b1: balance of sender, b2: balance of receiver}
	prev *Block
}
//Node: "Miners"
type Node struct {
	//Could store blocks in a map or in a linked list -
	name    int
	current Block
}

type log struct {
	nodes []Node
	lastValid Block
}

var logToNode map[int]chan Block
var nodeToLog chan Block
var logger log

//Checks validity of block
func checkValid(b Block, difficulty int) bool {
	sum := sha256.New()
	sum.Write([]byte(strconv.Itoa(b.name) + b.hash + strconv.Itoa(b.nonce)))
	sha1Hash := hex.EncodeToString(sum.Sum(nil))
	fmt.Println(sha1Hash)
	prefix := strings.Repeat("0", difficulty)
	if strings.HasPrefix(sha1Hash, prefix) {
		return true
	}
	return false
}

func (n *Node) mine(difficulty int) {
	oldHead := n.current
	transaction := fmt.Sprintf("%d=10", n.name)
	sum := sha256.New()
	sum.Write([]byte(transaction))
	hashTransaction := hex.EncodeToString(sum.Sum(nil))
	nonce := 1
	newBlock := Block{n.current.name + 1, nonce, hashTransaction, &oldHead}
	for {
		if checkValid(newBlock, difficulty) == true {
			break
		} else if oldHead != n.current {
			return
		}
		nonce = rand.Intn(MaxInt)
		newBlock.nonce = nonce
	}
	//TODO: pass valid block through channel to leader
	nodeToLog <- newBlock
}

//Simulate a node "A" checking validity of block again.
func loggerCheck(difficulty int, newBlock Block) bool{

	return true
}

func (n *Node) listen(difficulty int) {
	//while loop listens to channel
	//update current block in struct

	select {
	case newBlock, ok := <- nodeToLog:
		if ok {
			if loggerCheck(difficulty, newBlock) {
				updateNodes(newBlock)
			}
			break
		}
	default:
		break //handles case where channel is empty
	}
}

func protocol() {
	//node n1
	for i := range logger.nodes {
		//	go logger.nodes[i].mine()
		//	go logger.nodes[i].listen()
	}
}

//

//Updates nodes when a new valid block is added
func updateNodes( newBlock Block) bool{
	logger.lastValid = newBlock
	for i := range logger.nodes {
		logToNode[i] <- newBlock
		logger.nodes[i].current = <-logToNode[i]
	}

	//If success return true
}

func init() {
	numMiners := 4
	initBlock := Block{0, 0, "", nil}
	logger.lastValid = initBlock
	for i := 0; i < numMiners; i++ {
		logger.nodes = append(logger.nodes, Node{i, initBlock})
	}
}

func main() {
	checkValid(logger.current, 2)
}
