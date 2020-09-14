package main

import (
	ds "blockchain/datastructures"
	"fmt"
)

func main() {
	block1 := ds.NewBlock(1, "0(Genesis Block)", "value")
	block2 := ds.NewBlock(1, block1.Header.Hash, "value")

	blockchain := ds.NewBlockChain()
	blockchain.Insert(block1)
	blockchain.Insert(block2)

	// fmt.Println(blockchain.Get(2))
	str := blockchain.EncodeToJSON()

	blockchain.DecodeFromJSON([]byte(str))
	fmt.Println(blockchain.EncodeToJSON())
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
