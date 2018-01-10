package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
)

//Block represents an element within our Blockchain
type Block struct {
	Hash          string
	PrevBlockHash string
	Data          interface{}
}

func (b *Block) setHash() {
	hash := sha256.Sum256([]byte(b.PrevBlockHash + b.Data.(string)))
	b.Hash = hex.EncodeToString(hash[:])
}

//Blockchain represents our blockchain
type Blockchain struct {
	Blocks []*Block
}

//AddBlock Add a new block data to blockchain
func (bc *Blockchain) AddBlock(data interface{}) (newBlock *Block) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock = NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
	return
}

//NewBlock inserts a new block into our Blockchain
func NewBlock(data interface{}, previousBlock string) (block *Block) {
	block = &Block{
		Data:          data,
		PrevBlockHash: previousBlock,
	}
	block.setHash()
	return
}

//GenesisBlock creates the first block
func GenesisBlock() (bigBangBlock *Block) {
	return NewBlock("Big Bang! World was created!", "")
}

//NewBlockchain creates the blockchain
func NewBlockchain() (ourBlockchain *Blockchain) {
	return &Blockchain{[]*Block{GenesisBlock()}}
}
