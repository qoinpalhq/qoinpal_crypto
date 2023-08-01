package main

import (
	"fmt"

	ap "github.com/ayowilfred95/qoinpal_crypto/api"
	btc "github.com/ayowilfred95/qoinpal_crypto/bitcoin"
	"github.com/ayowilfred95/qoinpal_crypto/db"
	eth "github.com/ayowilfred95/qoinpal_crypto/ethereum"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	// got
	engine := gin.Default()
	fmt.Println("Multichain")
	fmt.Println(btc.NewBitcoinDisposableWallet())
	fmt.Println(eth.GenerateNewWallet())
	server := ap.NewApiServer(engine)
	server.RunServer()


}
