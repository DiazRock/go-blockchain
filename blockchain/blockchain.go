package blockchain

import (
	blocks "github.com/DiazRock/go-blockchain/block"
)

type Blockchain struct {
	Blocks []*blocks.Block
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := blocks.NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

func NewGenesisBlock() *blocks.Block {
	return blocks.NewBlock("Genesis Block", []byte{})
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*blocks.Block{NewGenesisBlock()}}
}
