package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

const MaxInt = int(^uint(0) >> 1)

//Block: Structure for holding Block information
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

//Channel for sending Blocks from miners to the logger
var nodeToLog chan Block

//Checks validity of the hash in block b for a certain difficulty
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

//init: Initializing Logger and miner nodes.
func init() {
	numMiners = 4
	difficulty = 5
	nodeToLog = make(chan Block, numMiners)
	initBlock := Block{0, 0, "", "", nil}
	logger.nodes = make(map[int]*Node)
	logger.lastValid = initBlock
	logger.newBlockChan = nodeToLog
	//Inits nodes in the logger
	for i := 0; i < numMiners; i++ {
		logToNode := make(chan Block)
		logger.nodes[i] = &Node{i, initBlock, logToNode}
	}
}

//runRounds: Runs protocol for a certain amount of rounds and keeps track of timing of each run
func runRounds(rounds int) map[int]string {
	//Time map to keep track of timing each round
	times := make(map[int]string)
	for i := 0; i < rounds; i++ {
		start := time.Now()
		protocol()
		end := time.Since(start)
		times[i] = fmt.Sprintf("%f", end.Seconds())
	}
	return times
}

func main() {
	t := runRounds(3)
	//Print time per round
	for k := range t {
		println(t[k])
	}
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
