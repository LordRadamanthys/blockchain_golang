package model

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"strconv"
)

type Block struct {
	Pos       int
	Data      BookCheckout
	TimeStamp string
	Hash      string
	PrevHash  string
}

func (b *Block) GenerateHash() {
	bytes, _ := json.Marshal(b.Data)

	data := strconv.Itoa(b.Pos) + b.TimeStamp + string(bytes) + b.PrevHash
	hs := sha256.New()
	hs.Write([]byte(data))
	b.Hash = hex.EncodeToString(hs.Sum(nil))
}

func (b *Block) ValidateBlock(block, prevBlock *Block) bool {
	if prevBlock.Hash != block.PrevHash {
		return false
	}

	if !block.ValidateHash(block.Hash) {
		return false
	}

	if prevBlock.Pos+1 != block.Pos {
		return false
	}
	return true
}

func (b *Block) ValidateHash(hash string) bool {
	b.GenerateHash()
	if b.Hash != hash {
		return false
	}
	return true
}
