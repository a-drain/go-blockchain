package main

import (
	ds "blockchain/datastructures"
	"fmt"
)

func main() {
	block1, err := ds.NewBlock(1, "0(Genesis Block)", "value")
	handleError(err)
	block2, err := ds.NewBlock(1, block1.Header.Hash, "value")
	handleError(err)

	blockchain := ds.NewBlockChain()
	err = blockchain.Insert(block1)
	handleError(err)

	err = blockchain.Insert(block2)
	handleError(err)

	// fmt.Println(blockchain.Get(2))
	str, err := blockchain.EncodeToJSON()
	handleError(err)

	blockchain.DecodeFromJSON([]byte(str))
	str, err = blockchain.EncodeToJSON()
	handleError(err)

	fmt.Println("Blockhain: ", str)
}
func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
