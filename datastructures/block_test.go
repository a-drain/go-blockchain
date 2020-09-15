package datastructures

import (
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
