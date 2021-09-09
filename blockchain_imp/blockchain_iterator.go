package blockchain_imp

import (
	"log"

	blocks "github.com/DiazRock/go-blockchain/block"
	"github.com/boltdb/bolt"
)

type BlockchainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

func (bc *Blockchain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{bc.Tip, bc.Db}

	return bci
}

func (i *BlockchainIterator) Next() *blocks.Block {
	var block *blocks.Block

	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(i.currentHash)
		block = blocks.DeserializeBlock(encodedBlock)

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
	i.currentHash = block.PrevBlockHash

	return block
}
