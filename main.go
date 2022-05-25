package main

import (
	"encoding/json"
	"fmt"

	"github.com/LordRadamanthys/blockchain_golang/src/adapters/input/controller"
	"github.com/LordRadamanthys/blockchain_golang/src/application/model"
	"github.com/LordRadamanthys/blockchain_golang/src/server"
)

func main() {
	Blockchain := InitBlockChain()

	blockChainController := controller.NewBlockchainController(Blockchain)
	routes := server.NewSatartApplication(blockChainController)

	routes.StartApplication()

	for _, block := range Blockchain.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevHash)
		bytes, _ := json.MarshalIndent(block.Data, "", "")
		fmt.Printf("Data: %v\n", string(bytes))
		fmt.Printf("Hash: %v\n", string(block.Hash))

	}
}

func InitBlockChain() *model.BlockChain {

	return &model.BlockChain{[]model.Block{*GenesisBlock()}}
}

func GenesisBlock() *model.Block {
	return model.CreateBlock(&model.Block{}, &model.BookCheckout{IsGenesis: true})
}
