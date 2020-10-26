package datastructures

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBlockchain(t *testing.T) {
	bGot := NewBlockChain()
	assert.NotNil(t, bGot, "New blockchain is nil")
	assert.True(t, bGot.Length == 0, "Length should be equal to zero upon new blockchain")
}

func TestInsert(t *testing.T) {
	b := NewBlockChain()
	block1, _ := NewBlock(1, "parent", "block1")
	err := b.Insert(block1)

	assert.True(t, err == nil, err)
	assert.EqualValues(t, b.Chain[1][0], block1, "Single fork, single block insertion failed")

	block2, _ := NewBlock(1, block1.Header.Hash, "block2")
	err = b.Insert(block2)
	assert.EqualValues(t, b.Chain[1][1], block2, "Single fork, more than one block insertion failed")

	block3, _ := NewBlock(2, block2.Header.Hash, "block3")
	err = b.Insert(block3)
	assert.EqualValues(t, b.Chain[2][0], block3, "Second fork, single value insertion failed")

	block4, _ := NewBlock(2, block3.Header.Hash, "block4")
	err = b.Insert(block4)
	assert.EqualValues(t, b.Chain[2][1], block4, "Second fork, more than one block insertion failed")

	err = b.Insert(block4)
	assert.Error(t, err, "Should not insert block with same hash")

}

func TestGet(t *testing.T) {
	b := NewBlockChain()

	block1, _ := NewBlock(1, "parent", "block1")
	block2, _ := NewBlock(1, block1.Header.Hash, "block2")

	b.Insert(block1)
	b.Insert(block2)

	rec, err := b.Get(1)
	exp := []*Block{block1, block2}

	assert.NoError(t, err, "Get threw an error")
	assert.EqualValues(t, exp, rec, "Should get blocks from the specified height")

	rec, err = b.Get(3)
	assert.Error(t, err, "Should return error upon get to height not in chain")

}

func TestEncodeBlockchainToJSON(t *testing.T) {
	b := NewBlockChain()
	block1, _ := NewBlock(1, "parent", "block1")
	block2, _ := NewBlock(1, block1.Header.Hash, "block2")
	b.Insert(block1)
	b.Insert(block2)

	exp := fmt.Sprintf(`{"header":{"height":1,"timestamp":%v,"hash":"%v","parentHash":"%v","size":32},"value":"%v"}{"header":{"height":1,"timestamp":%v,"hash":"%v","parentHash":"%v","size":32},"value":"%v"}`,
		block1.Header.Timestamp, block1.Header.Hash, block1.Header.ParentHash, block1.Value, block2.Header.Timestamp, block2.Header.Hash, block2.Header.ParentHash, block2.Value)
	rec, err := b.EncodeToJSON()

	assert.NoError(t, err, "Encoding threw an error when it wasn't supposed to")
	assert.Equal(t, exp, rec, "Expected json does not match received json")
}

func TestBlockchainDecodeFromJSON(t *testing.T) {
	block1, _ := NewBlock(1, "parent", "block1")
	block2, _ := NewBlock(1, block1.Header.Hash, "block2")
	jsonBlockchain := fmt.Sprintf(`[{"header":{"height":1,"timestamp":%v,"hash":"%v","parentHash":"%v","size":32},"value":"%v"},{"header":{"height":1,"timestamp":%v,"hash":"%v","parentHash":"%v","size":32},"value":"%v"}]`,
		block1.Header.Timestamp, block1.Header.Hash, block1.Header.ParentHash, block1.Value, block2.Header.Timestamp, block2.Header.Hash, block2.Header.ParentHash, block2.Value)

	b := NewBlockChain()
	err := b.DecodeFromJSON(nil)
	assert.Error(t, err, "Should throw error upon bad input")

	err = b.DecodeFromJSON([]byte(jsonBlockchain))
	expChain := NewBlockChain()
	expChain.Insert(block1)
	expChain.Insert(block2)
	assert.EqualValues(t, expChain, b, "Should add blocks to new blockchain")

}
