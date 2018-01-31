package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"strconv"
	"time"

	"github.com/novatrixtech/tonecoin/proto"
)

//Block represents an element within our Blockchain
type Block struct {
	Timestamp     int64
	Hash          string
	PrevBlockHash string
	Data          interface{}
	Datatype      proto.Datatype
}

func (b *Block) setHash() {
	hash := sha256.Sum256([]byte(b.PrevBlockHash + b.Data.(string) + strconv.Itoa(int(b.Timestamp))))
	b.Hash = hex.EncodeToString(hash[:])
}

//Blockchain represents our blockchain
type Blockchain struct {
	Blocks []*Block
}

//AddBlock Add a new block data to blockchain
func (bc *Blockchain) AddBlock(data interface{}, datatype proto.Datatype) (newBlock *Block) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock = NewBlock(data, datatype, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
	return
}

//NewBlock inserts a new block into our Blockchain
func NewBlock(data interface{}, datatype proto.Datatype, previousBlock string) (block *Block) {
	block = &Block{
		Timestamp:     time.Now().Unix(),
		Data:          data,
		PrevBlockHash: previousBlock,
		Datatype:      datatype,
	}
	block.setHash()
	log.Printf("Block criado é: %+v \r\n", block)
	return
}

//GenesisBlock creates the first block
func GenesisBlock() (bigBangBlock *Block) {
	return NewBlock("Big Bang! World was created!", proto.Datatype_TEXT, "")
}

//NewBlockchain creates the blockchain
func NewBlockchain() (ourBlockchain *Blockchain) {
	return &Blockchain{[]*Block{GenesisBlock()}}
}
