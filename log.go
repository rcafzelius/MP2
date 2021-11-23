package main

import "sync"

type Log struct {
	nodes        map[int]*Node
	lastValid    Block
	newBlockChan chan Block
}

var mu sync.Mutex

//Check validity of newBlock hash
func loggerCheck(difficulty int, newBlock Block) bool {
	return checkValid(newBlock, difficulty)
}

//Contact nodes when new valid block has been added
func (l *Log) updateNodes(newBlock Block) {
	l.lastValid = newBlock
	for k := range l.nodes {
		l.nodes[k].logToNode <- l.lastValid
	}
}

//Check the nodeToLog channel and call updateNodes if newBlock's hash is valid
func (l *Log) checkForBlock() {
L:
	for {
		select {
		case newBlock, ok := <-l.newBlockChan:
			if ok {
				mu.Lock()
				if loggerCheck(difficulty, newBlock) {
					l.updateNodes(newBlock)
					mu.Unlock()
					break L
				}
				//Unlocks nodeToLog channel
				mu.Unlock()
			}
		default:
		}
	}
}

//Clears the nodeToLog channel
func (l *Log) clearChannel() {
	for i := 0; i < numMiners; i++ { //since we will set all nodes' channels to 0's, only need to clear 0's channels
		select {
		case _, ok := <-l.newBlockChan:
			if ok {
			}
		default:
			break //handles case where channel is empty
		}
	}
}
