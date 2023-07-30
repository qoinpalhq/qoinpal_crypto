package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"strings"
	"net/http"
	btc "github.com/ayowilfred95/qoinpal_crypto/bitcoin"
	eth "github.com/ayowilfred95/qoinpal_crypto/ethereum" 
)


func init(){
	log.SetPrefix("API:")
}

type ApiServer struct{
	Router *gin.Engine
}

func NewApiServer(router *gin.Engine) *ApiServer{
	return &ApiServer{router}
}
func (ap *ApiServer) RunServer(){
	// register endpoints
	ap.Router.GET("/api/address/:chain_type", ap.handleGetCryptoAddress)
	ap.Router.Run()
}

func (ap *ApiServer) handleGetCryptoAddress(c *gin.Context) {
	chainType := c.Param("chain_type")
	chainType = strings.ToLower(chainType)

	switch chainType {
	case "bitcoin":
		newWallet, err := btc.NewBitcoinDisposableWallet()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to generate Bitcoin wallet",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"bitcoin_address": newWallet.Address,
		})

	case "ethereum":
		newWallet, err := eth.GenerateNewWallet()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to generate Ethereum wallet",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"ethereum_address": newWallet.Address,
		})

	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid chain type",
		})
	}
}
