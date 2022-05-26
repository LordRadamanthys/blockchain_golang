package model

import (
	"time"
)

type BlockChain struct {
	Blocks []*Block
}

func CreateBlock(prevBlock *Block, checkoutItem BookCheckout) *Block {
	block := &Block{}
	block.Pos = prevBlock.Pos + 1
	block.TimeStamp = time.Now().String()
	block.PrevHash = prevBlock.Hash
	block.Data = checkoutItem
	block.GenerateHash()

	return block
}
func (bc *BlockChain) AddBlock(bookCheckout BookCheckout) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	block := CreateBlock(prevBlock, bookCheckout)

	if block.ValidateBlock(block, prevBlock) {
		bc.Blocks = append(bc.Blocks, block)
	}
}
