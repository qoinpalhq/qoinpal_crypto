package main

import ( 
	"fmt"
	btc "github.com/ayowilfred95/qoinpal_crypto/bitcoin"
	ap "github.com/ayowilfred95/qoinpal_crypto/api"
	"github.com/gin-gonic/gin"

)

func main(){
	// got
	engine := gin.Default()
	fmt.Println("Multichain")
	fmt.Println(btc.NewBitcoinDisposableWallet())
	server :=ap.NewApiServer(engine)
	server.RunServer()

}