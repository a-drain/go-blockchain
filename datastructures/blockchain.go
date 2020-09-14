package datastructures

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Blockchain is the blockchain datastructure
type Blockchain struct {
	Chain  map[int32][]*Block `json:"chain"`
	Length int32              `json:"length"`
}

// NewBlockChain initializes a new blockchain struct
func NewBlockChain() *Blockchain {
	var b Blockchain
	b.Chain = make(map[int32][]*Block)
	return &b
}

// Get used to get blocks of a specified height
func (b *Blockchain) Get(height int32) []*Block {
	if val, ok := b.Chain[height]; ok {
		return val
	}
	return nil
}

// Insert inserts a block into the blockchain
func (b *Blockchain) Insert(block *Block) {

	if len(b.Chain[block.Header.Height]) == 0 {
		var newSlice []*Block = []*Block{block}
		b.Chain[block.Header.Height] = make([]*Block, 2)
		b.Chain[block.Header.Height] = newSlice
		return
	}

	for i := range b.Chain[block.Header.Height] {
		if b.Chain[block.Header.Height][i].Header.Hash != block.Header.Hash {
			b.Chain[block.Header.Height] = append(b.Chain[block.Header.Height], block)
			return
		}
	}
}

// EncodeToJSON encodes a blockchain to json
func (b *Blockchain) EncodeToJSON() string {
	var str strings.Builder

	for _, v := range b.Chain {
		for _, val := range v {
			json, err := val.EncodeToJSON()
			if err != nil {
				panic(err)
			}
			_, err = str.Write(json)
			if err != nil {
				panic(err)
			}
		}
	}
	return str.String()
}

// DecodeFromJSON decodes a string to a list of block JSON strings and inserts each block into the chain
func (b *Blockchain) DecodeFromJSON(jsonData []byte) error {
	tmp := []interface{}{&b.Chain}
	wantLen := len(tmp)

	if err := json.Unmarshal(jsonData, &tmp); err != nil {
		return err
	}
	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("Wrong number of fields in Blockchain: %d != %d", g, e)
	}
	return nil
}
