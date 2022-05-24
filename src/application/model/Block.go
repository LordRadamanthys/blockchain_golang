package model

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
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

	data := string(b.Pos) + b.TimeStamp + string(bytes) + b.PrevHash
	hs := sha256.New()
	hs.Write([]byte(data))
	b.Hash = hex.EncodeToString(hs.Sum(nil))
}
