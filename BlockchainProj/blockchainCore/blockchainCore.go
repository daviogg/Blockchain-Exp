package blockchaincore

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

//Block - struct
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

//Blockchain - struct
type Blockchain struct {
	blocks []*Block
}

//GetBlocks - return blocks in blockchain object
func (bc Blockchain) GetBlocks() []*Block {
	return bc.blocks
}

//SetHash - Hash block with sha256
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

//NewBlock - create new block and hash them
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

//AddBlock - add block to blochain blocks
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

//NewGenesisBlock - create a genesis block for validate first addBlock
func NewGenesisBlock() *Block {
	return NewBlock("DavioGG genesis Block", []byte{})
}

//NewBlockchain - Create a new istance of blockchain
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
