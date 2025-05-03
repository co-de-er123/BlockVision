package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"your_project/handlers"
)

func main() {
	r := gin.Default()
	r.GET("/block/:blockNumber", handlers.GetBlockDetails)
	r.GET("/transaction/:txHash", handlers.GetTransactionDetails)
	r.GET("/network-health", handlers.GetNetworkHealth)

	err := r.Run(":8080")
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
