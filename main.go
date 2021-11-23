package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

const MaxInt = int(^uint(0) >> 1)

type Block struct {
	name, nonce int
	hash        string
	transaction string
	//add single transaction as map: {source: location coin came from, b1: balance of sender, b2: balance of receiver}
	prev *Block
}

//Node: "Miners"
var logger Log
var difficulty int
var numMiners int
var nodeToLog chan Block

//Checks validity of block
func checkValid(b Block, difficulty int) bool {
	sum := sha256.New()
	sum.Write([]byte(strconv.Itoa(b.name) + b.hash + strconv.Itoa(b.nonce)))
	sha1Hash := hex.EncodeToString(sum.Sum(nil))
	prefix := strings.Repeat("0", difficulty)
	if strings.HasPrefix(sha1Hash, prefix) {
		return true
	}
	return false
}

func protocol() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		logger.checkForBlock()
	}()
	for i := 0; i < numMiners; i++ {
		n := logger.nodes[i]
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			n.mine(difficulty)
		}(i)
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			n.listen()
		}(i)
	}
	wg.Wait()
	logger.clearChannel()
	fmt.Println("One round done")
}

func init() {
	numMiners = 4
	difficulty = 5
	nodeToLog = make(chan Block, numMiners)
	initBlock := Block{0, 0, "", "", nil}
	logger.nodes = make(map[int]*Node)
	logger.lastValid = initBlock
	logger.newBlockChan = nodeToLog
	for i := 0; i < numMiners; i++ {
		logToNode := make(chan Block)
		logger.nodes[i] = &Node{i, initBlock, logToNode}
	}
}

func runRounds(rounds int) {
	for i := 0; i < rounds; i++ {
		v := logger.nodes[0].current
		print(v.name)
		protocol()
	}
}

func main() {
	runRounds(3)
	a := Block{}
	b := logger.lastValid
	for {
		print(b.name)
		b = *b.prev
		if b == a {
			break
		}
	}

}
