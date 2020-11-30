package api

import "github.com/gin-gonic/gin"

func SetupTransactionAPI(router *gin.Engine) {
	transactionAPI := router.Group("/api/v2")
	{
		transactionAPI.POST("/send", send)
		//productAPI.POST("/create", Create)
	}
}

// Send - send api
func send(c *gin.Context) {
	c.JSON(401, gin.H{"result": "Send"})
}
