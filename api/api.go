package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"strings"
	"net/http"
	btc "github.com/ayowilfred95/qoinpal_crypto/bitcoin"
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

func (ap *ApiServer) handleGetCryptoAddress(c *gin.Context){
	chainType := c.Param("chain_type")
	// if len(chainType) == 0{
		// c.JSON(http.StatusBadRequest, gin.H{
			// "error":"invalid chain type, value cannot be blank",
		// })
	// }
	// convert type to all lower
	chainType = strings.ToLower(chainType)

	if chainType == "bitcoin"{
		newWallet := btc.NewBitcoinDisposableWallet()
		c.JSON(http.StatusOK, gin.H{
			"bitcoin_address": newWallet.Address,
		})
	}
}