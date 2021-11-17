package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

type Block struct {
	name, hash, nonce int
	//add single transaction as map: {source: location coin came from, b1: balance of sender, b2: balance of receiver}
	prev *Block
}
type Node struct {
	//Could store blocks in a map or in a linked list -
	name    int
	current Block
}

type log struct {
	nodes   []Node
	current Block
}

var logToNode map[int]chan Block
var nodeToLog chan Block
var logger log

//Checks validity of block
func checkValid(b Block, difficulty int) bool {
	sum := sha256.New()
	sum.Write([]byte(strconv.Itoa(b.name) + strconv.Itoa(b.hash) + strconv.Itoa(b.nonce)))
	sha1Hash := hex.EncodeToString(sum.Sum(nil))
	fmt.Println(sha1Hash)
	prefix := strings.Repeat("0", difficulty)
	if strings.HasPrefix(sha1Hash, prefix) {
		return true
	}
	return false
}

func (*Node) mine() {
	//assemble first try
	//create transaction
	//Block tmp
	//while checkValid(block) is false:
	//		if tmp.current != Node.current: return //terminate mining step if someone else mines the block first
	//		assemble new block with random nonce
	//		new block has previous = current, nonce= randomly generated, hash = ?, name = node.name

	//pass valid block through channel to leader
}

func (*Node) listen() {
	//while loop listens to channel
	//update current block in struct
	/*
		x, err <- channel
		err = nil
		for {
			x, err <- channel
			if err != nil{
				continue
			}
		}
	*/
}

func protocol() {
	//node n1
	for i := 0; i < 5; i++ {
		//	go n1.mine()
		//	go n1.listen()
	}
}

//

//Updates nodes when a new valid block is added
func updateNodes() {

}

func init() {
	numMiners := 4
	initBlock := Block{0, 0, 0, nil}
	logger.current = initBlock
	for i := 0; i < numMiners; i++ {
		logger.nodes = append(logger.nodes, Node{i, initBlock})
	}
}

func main() {
	checkValid(logger.current, 2)
}
