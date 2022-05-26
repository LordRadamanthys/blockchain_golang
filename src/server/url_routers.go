package server

import (
	"github.com/LordRadamanthys/blockchain_golang/src/application/ports"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

type BController struct {
	BlockchainController ports.BlockchainController
}

func NewSatartApplication(bc ports.BlockchainController) *BController {
	return &BController{
		BlockchainController: bc,
	}
}

func (bc *BController) StartApplication() {

	router.GET("/blockchain", bc.BlockchainController.GetBlockchain)
	router.POST("/blockchain", bc.BlockchainController.WriteBlockchain)
	router.POST("/book", bc.BlockchainController.NewBook)

	router.Run(":8080")
}
