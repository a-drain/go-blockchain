package datastructures

import (
	"encoding/json"
	"errors"
	"strings"
)

// Blockchain is the blockchain datastructure
type Blockchain struct {
	Chain  map[int32][]*Block `json:"chain"`
	Length int32              `json:"length"`
}

// NewBlockChain initializes a new blockchain struct
func NewBlockChain() *Blockchain {
	return &Blockchain{Chain: make(map[int32][]*Block), Length: 0}
}

// Get used to get blocks of a specified height
func (b *Blockchain) Get(height int32) ([]*Block, error) {
	if val, ok := b.Chain[height]; ok {
		return val, nil
	}
	return nil, errors.New("Block does not exist at provided height")
}

// Insert inserts a block into the blockchain
func (b *Blockchain) Insert(block *Block) error {

	if len(b.Chain[block.Header.Height]) == 0 {
		b.Chain[block.Header.Height] = make([]*Block, 1)
		b.Chain[block.Header.Height][0] = block
		b.Length++
		return nil
	}

	slice := b.Chain[block.Header.Height]

	for i := range slice {
		if b.Chain[block.Header.Height][i].Header.Hash != block.Header.Hash {
			slice = append(slice, block)
			b.Chain[block.Header.Height] = slice
			b.Length++
		} else {
			return errors.New("Cannot insert, block with same hash already exists")
		}
	}
	return nil
}

// EncodeToJSON encodes a blockchain to json
func (b *Blockchain) EncodeToJSON() (string, error) {
	var str strings.Builder

	for _, fork := range b.Chain {
		for _, block := range fork {
			json, err := block.EncodeToJSON()
			if err != nil {
				return "", err
			}
			_, err = str.Write(json)
			if err != nil {
				return "", err
			}
		}
	}
	return str.String(), nil
}

// DecodeFromJSON decodes a string to a list of block JSON strings and inserts each block into the chain
func (b *Blockchain) DecodeFromJSON(jsonData []byte) error {
	var blocks []*Block
	if err := json.Unmarshal(jsonData, &blocks); err != nil {
		return err
	}
	for _, block := range blocks {
		if err := b.Insert(block); err != nil {
			return err
		}
	}
	return nil
}
