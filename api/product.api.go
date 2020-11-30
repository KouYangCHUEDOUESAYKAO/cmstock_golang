package api

import "github.com/gin-gonic/gin"

func SetupProductAPI(router *gin.Engine) {
	productAPI := router.Group("/api/v2")
	{
		productAPI.POST("/list", list)
		productAPI.POST("/create", create)
	}
}

// Llis - list api
func list(c *gin.Context) {
	c.JSON(401, gin.H{"result": "list"})
}

// Create - create api
func create(c *gin.Context) {
	c.JSON(401, gin.H{"result": "create"})
}
