package blockchain

//Block represents an element within our Blockchain
type Block struct {
	Hash          string
	PrevBlockHash string
	Data          interface{}
}

//Blockchain represents our blockchain
type Blockchain struct {
	Blocks []*Block
}

//NewBlock inserts a new block into our Blockchain
func NewBlock(data interface{}, previousBlock string) (block *Block) {
	block = &Block{
		Data:          data,
		PrevBlockHash: previousBlock,
	}
	return
}
