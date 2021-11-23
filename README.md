# MP2
Simulates a blockchain through the use of a lead node and mining nodes that communicate through the leader.

## How to Run
1. Download the repository and navigate to the folder with the go files. To modify aspects of the protocol, change the following:
* numMiners = the number of mining nodes used, in init()
* difficulty = the number of leading zeros necessary for a block to solve the complexity puzzle, in init()
* runRounds(rounds) takes an integer representing how many blocks will be created, in main()
* runtime.GOMAXPROCS(i) takes an integer representing the number of operating systems threads that can be used, in main()
2. Enter `go run .`.
3. By default, five blocks are created on the chain, and the function outputs how long each block took to create.

## Design
### Main
Globally, we stored a Log struct object, a channel for mining nodes to send completed blocks to the leader, and
integer values associated with the number of mining nodes and the puzzle difficulty. Puzzles were completed in
a similar manner to bitcoin, where a hash of the transaction data, the name of the new block, and a nonce value
were concatenated and hashed with sha256, which was then checked to confirm it has at least the number of leading zeros 
specified by the difficulty value. This was done by turning the hash value (in hex form) into a string, creating a string 
of zeros equal to the number specified by the difficulty, and checking if the hash contained the string of zeros as a
prefix (the first characters in the string). This is imperfect, as the value 0x0 is equivalent to b'0000, meaning our
difficulty value actually adds four bits of difficulty each time it is incremented, but it simplifies the storage and check
prefix check steps, as the values in question are significantly smaller. `Protocol()` runs the logger's listening function,
as well as each mining node's listening and mining functions as independent gorountines. For four mining nodes, this 
translates to 9 goroutines. We use a waitgroup that waits for all these routines to terminate (how this is done is discussed
below, then we print that a round has completed). `RunRounds()` runs the protocol function a specified number of times,
and times each run and stores it in a map, which is returned after the last round completes. `init()` initializes the logger,
sets values for global variables, and creates the 'nodeToLog' channel, the channel that nodes use to send completed nodes to 
the logger.
### Log
The lead node, called 'Log' or 'logger', stores a map of pointers to mining nodes, the last valid block created, and a
channel for nodes to send blocks to it. The primary functionality of the logger is to receive nodes and check their validity
before distributing the new block. `checkForBlock()` listens to the channel stored in logger. When a block is received, 
the function locks and checks the validity of the block. This is to prevent a conflict where a second block is considered
valid before the first block is distributed. `loggerCheck()` checks the validity of the new block using the
`checkValid()` function stored in main. If the new block is valid, it is passed to `updateNodes()`, which stores the new
block locally, then distributes it via the map of pointers to nodes. The specifics of this mechanism are discussed below. 
Once a round has been completed, the protocol function in main calls `clearChannel`, which pulls values from the 
newBlockChan channel until it is emptied. This is to ensure that blocks working on older parts of the chain are disregarded.
### Node
Nodes store a unique name, a copy of the current block, and a channel for sending blocks to the logger. When the logger
is distributing new blocks, it does so by accessing this channel via its personal map of nodes and sending the block. 
`listen()` listens to this personal channel, and if a new block is received, checks if the new block is equal to the one
currently stored (if it is, the node must have mined it; if it isn't, this value is updated), then terminates. On the mining
side, we implemented two situations simultaneously: one where new blocks are created to be valid, and one where randomly,
a junk node is created. This was done to confirm the logger could handle invalid blocks and not add them to the chain. 
Mining is done by creating a new block that points to the current block, has a name as an integer 1+the name of the current
block, a transaction, a sha256 hash of that transaction, and a nonce originally equal to 1. To keep things simple, all our 
transactions assign 10 units (i.e. BTC) to whoever created the block. The function then loops with two conditions for 
termination: if the node successfully creates a block, in which case it is sent to the logger via the nodeToLog channel,
or if another node creates a block, which is detected when the former current node and the one the node stores in its 
current field are no longer equal. Otherwise, the nonce is recomputed, in our case done as one of all positive integers 
up to the maximum value that can be stored (see const MaxInt in main). Randomly, if this value is less than the ratio provided,
then the block, which may or may not be valid is sent to the logger, a process that simulates faulty behavior on the part
of the nodes. Since this process is random, we do not terminate mining as a punishment. As mentioned above, both mine and 
listen are run as goroutines, and therefore both must terminate for the protocol function to operate as specified. Mine 
terminates when either a valid block is generated or a new block is stored in the node (as a result of listen), and listen 
terminates when a new block is transmitted, so this ensures the waitgroup in protocol does not wait endlessly. 

#### Sources
Checking zeros as a prefix of a hash value: https://mycoralhealth.medium.com/code-your-own-blockchain-mining-algorithm-in-go-82c6a71aba1f