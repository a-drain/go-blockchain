package datastructures

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestNewBlock tests the creation of a new block
func TestNewBlock(t *testing.T) {
	block, err := NewBlock(2, "genesis", "test")

	assert.True(t, err == nil, "Asserting error is not thrown from creation of block.New")
	assert.True(t, block.Header.Height == 2, "Asserting height is correct")
	assert.True(t, block.Header.ParentHash == "genesis", "Asserting parent hash is correct")

	expDay := time.Unix(block.Header.Timestamp, 0).Day()
	currDay := time.Now().Day()
	assert.True(t, expDay == currDay, "Asserting timestamp created is correct day")

}

func TestGenHash(t *testing.T) {
	hash1, err1 := GenHash(2, "parent", "value")
	time.Sleep(1 * time.Second)
	hash2, err2 := GenHash(2, "parent", "value")
	assert.True(t, err1 == nil || err2 == nil, "Asserting error is not thrown from hash creation")
	assert.True(t, hash1 != hash2, "Asserting hashes are not the same")
}

func TestDecodeFromJSON(t *testing.T) {
	tm := time.Now().Unix()
	jsonStr := fmt.Sprintf(`{"header":{"height": 0, "hash": "hash", "timestamp":  %d, "parentHash": "parentHash", "size": 32}, "value": "asdf"}`, tm)
	jsonBytes := []byte(jsonStr)

	gotBlock, err := DecodeFromJSON(jsonBytes)
	fmt.Println("Value", gotBlock.Value)

	expectedBlock := &Block{Header: Header{
		Height:     0,
		Hash:       "hash",
		Timestamp:  tm,
		ParentHash: "parentHash",
		Size:       32,
	},
		Value: "asdf",
	}

	assert.True(t, err == nil, fmt.Sprintf("Error recieved while unmarshalling: %v", err))
	assert.Equal(t, expectedBlock, gotBlock, "Block json unmarshalling gone wrong")
}

func TestEncodeToJSON(t *testing.T) {
	tm := time.Now().Unix()

	block := &Block{Header: Header{
		Height:     0,
		Hash:       "hash",
		Timestamp:  tm,
		ParentHash: "parentHash",
		Size:       32,
	},
		Value: "asdf",
	}

	gotJSON, err := block.EncodeToJSON()
	gotStr := string(gotJSON)
	expectedJSON := fmt.Sprintf(`{"header":{"height":0,"timestamp":%d,"hash":"hash","parentHash":"parentHash","size":32},"value":"asdf"}`, tm)

	assert.True(t, err == nil, "Asserting error is not nil")
	assert.Equal(t, expectedJSON, gotStr, "Block json marshalling gone wrong ")
}
