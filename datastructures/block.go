package datastructures

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"strconv"
	"time"
)

type (
	// Header is a block header
	Header struct {
		Height     int32  `json:"height"`
		Timestamp  int64  `json:"timestamp"`
		Hash       string `json:"hash"`
		ParentHash string `json:"parentHash"`
		Size       int32  `json:"size"`
	}

	// Block is a block in the blockchain
	Block struct {
		Header Header `json:"header"`
		Value  string `json:"value"`
	}
)

// NewBlock creates a block
func NewBlock(height int32, parentHash string, value string) *Block {
	hash, err := genHash(height, parentHash, value)
	if err != nil {
		panic(err)
	}
	return &Block{
		Header: Header{
			Height:     height,
			Hash:       hash,
			Timestamp:  time.Now().Unix(),
			ParentHash: parentHash,
			Size:       32,
		},
		Value: value,
	}
}

func genHash(height int32, parentHash string, value string) (string, error) {
	var unixTime int64 = time.Now().Unix()
	currentTime := strconv.FormatInt(unixTime, 10)
	hashStr := string(height) + currentTime + parentHash + string(32) + value
	hasher := sha1.New()
	if _, err := hasher.Write([]byte(hashStr)); err != nil {
		return "", err
	}

	sha := base64.StdEncoding.EncodeToString(hasher.Sum(nil))
	return sha, nil
}

// DecodeFromJSON to decode a json string to a block
func DecodeFromJSON(jsonData []byte) (Block, error) {
	block := Block{}
	err := json.Unmarshal(jsonData, &block)
	if err != nil {
		return block, err
	}
	return block, nil
}

// EncodeToJSON to encode a block to json
func (b Block) EncodeToJSON() ([]byte, error) {

	ToJSON, err := json.Marshal(b)
	if err != nil {
		return ToJSON, err
	}

	return ToJSON, nil
}
