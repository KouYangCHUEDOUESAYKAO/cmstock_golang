package api

import "github.com/gin-gonic/gin"

func SetupTransactionAPI(router *gin.Engine) {
	transactionAPI := router.Group("/api/v2")
	{
		transactionAPI.GET("/transaction", transaction)
		transactionAPI.POST("/transaction", createtransaction)
	}
}

// Transaction - transaction api
func transaction(c *gin.Context) {
	c.JSON(401, gin.H{"result": "List Transaction"})
}

// Createtranaction - createtransaction api
func createtransaction(c *gin.Context) {
	c.JSON(401, gin.H{"result": "Create Transaction"})
}
