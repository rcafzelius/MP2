package main

type Log struct {
	nodes        map[int]Node
	lastValid    Block
	newBlockChan chan Block
}

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
				if loggerCheck(difficulty, newBlock) {
					l.updateNodes(newBlock)
					break L
				}
			}
		default:
		}
	}
	println("sent")
}
