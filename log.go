package main

import "sync"

type Log struct {
	nodes        map[int]Node
	lastValid    Block
	newBlockChan chan Block
}

var mu sync.Mutex

func loggerCheck(difficulty int, newBlock Block) bool {
	return checkValid(newBlock, difficulty)
}

//Contact nodes when new valid block has been added
func (l *Log) updateNodes(newBlock Block) {
	l.lastValid = newBlock
	for k := range l.nodes {
		l.nodes[k].logToNode <- newBlock
	}
}

//
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
				mu.Unlock()
			}
		default:
		}
	}
	println("sent")
}
