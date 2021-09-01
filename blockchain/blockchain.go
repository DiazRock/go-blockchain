package blockchain

import (
	blocks "github.com/DiazRock/go-blockchain/blocks"
)

type Blockchain struct {
	blocks []*blocks.Block
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := blocks.NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

func NewGenesisBlock() *blocks.Block {
	return blocks.NewBlock("Genesis Block", []byte{})
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*blocks.Block{NewGenesisBlock()}}
}
