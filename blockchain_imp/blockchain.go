package blockchain_imp

import (
	"log"

	blocks "github.com/DiazRock/go-blockchain/block"
	"github.com/boltdb/bolt"
)

const (
	dbFile       = "blockchain.go"
	blocksBucket = "blocks"
)

type Blockchain struct {
	Tip []byte
	Db  *bolt.DB
}

func (bc *Blockchain) AddBlock(data string) {
	var lastHash []byte

	err := bc.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("l"))

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	newBlock := blocks.NewBlock(data, lastHash)

	err = bc.Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put(newBlock.Hash, newBlock.Serialize())

		if err != nil {
			log.Fatal(err)
		}

		err = b.Put([]byte("l"), newBlock.Hash)

		if err != nil {
			log.Fatal(err)
		}
		bc.Tip = newBlock.Hash

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func NewBlockchain() *Blockchain {
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		if b == nil {
			genesis := blocks.NewGenesisBlock()
			b, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				log.Panic(err)
			}
			err = b.Put(genesis.Hash, genesis.Serialize())
			if err != nil {
				log.Fatal(err)
			}
			err = b.Put([]byte("l"), genesis.Hash)
			if err != nil {
				log.Fatal(err)
			}
			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("l"))
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
	bc := Blockchain{tip, db}

	return &bc
}
