package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"your_project/services"
	"strconv"
)

func GetBlockDetails(c *gin.Context) {
	blockNumber, err := strconv.ParseUint(c.Param("blockNumber"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid block number"})
		return
	}

	block, err := services.GetBlockDetails(blockNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch block details"})
		return
	}

	c.JSON(http.StatusOK, block)
}

func GetTransactionDetails(c *gin.Context) {
	txHash := c.Param("txHash")
	tx, err := services.GetTransactionDetails(txHash)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transaction details"})
		return
	}

	c.JSON(http.StatusOK, tx)
}

func GetNetworkHealth(c *gin.Context) {
	health, err := services.GetNetworkHealth()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch network health"})
		return
	}

	c.JSON(http.StatusOK, health)
}
