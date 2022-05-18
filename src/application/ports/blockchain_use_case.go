package ports

import "github.com/gin-gonic/gin"

type BlockchainController interface {
	GetBlockchain(*gin.Context)
	WriteBlockchain(*gin.Context)
	NewBook(*gin.Context)
}
