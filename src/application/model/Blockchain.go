package model

import "time"

type BlockChain struct {
	Blocks []Block
}

func CreateBlock(prevBlock *Block, checkoutItem *BookCheckout) *Block {
	block := &Block{}
	block.Pos = prevBlock.Pos + 1
	block.TimeStamp = time.Now().String()
	block.PrevHash = prevBlock.Hash
	block.GenerateHash()

	return block
}
func (bc *BlockChain) AddBlock(bookCheckout *BookCheckout) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	block := CreateBlock(&prevBlock, bookCheckout)

	if validBlock(block, &prevBlock) {
		bc.Blocks = append(bc.Blocks, *block)
	}
}

func validBlock(block, prevBlock *Block) bool {
	if block.Hash != prevBlock.PrevHash {
		return false
	}

	if block.validateBlock(block.Hash) {
		return false
	}

	if prevBlock.Pos+1 != block.Pos {
		return false
	}

	return true
}
