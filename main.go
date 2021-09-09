package main

import (
	blockchain "github.com/DiazRock/go-blockchain/blockchain_imp"
	cli "github.com/DiazRock/go-blockchain/cli_imp"
)

func main() {
	Bc := blockchain.NewBlockchain()
	defer Bc.Db.Close()
	cli_instance := cli.CLI{Bc}
	cli_instance.Run()
}
