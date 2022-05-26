package controller

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"

	"github.com/LordRadamanthys/blockchain_golang/src/application/model"
	"github.com/gin-gonic/gin"
)

type BlockChainController struct {
	Blockchain *model.BlockChain
}

func NewBlockchainController(block *model.BlockChain) *BlockChainController {
	return &BlockChainController{
		Blockchain: block,
	}
}

func (bc *BlockChainController) NewBook(ctx *gin.Context) {

	var book model.Book

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(400, err)
		return
	}

	h := md5.New()
	io.WriteString(h, book.ISBN+book.PublishDate)
	book.ID = fmt.Sprintf("%x", h.Sum(nil))

	ctx.JSON(http.StatusCreated, book)
}

func (bc *BlockChainController) WriteBlockchain(ctx *gin.Context) {

	var checkoutItem model.BookCheckout

	if err := ctx.ShouldBindJSON(&checkoutItem); err != nil {
		ctx.JSON(400, err)
		return
	}

	bc.Blockchain.AddBlock(checkoutItem)
	ctx.JSON(http.StatusOK, checkoutItem)
}

func (bc *BlockChainController) GetBlockchain(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, bc.Blockchain)
}
