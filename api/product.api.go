package api

import "github.com/gin-gonic/gin"

func SetupProductAPI(router *gin.Engine) {
	productAPI := router.Group("/api/v2")
	{
		productAPI.GET("/product", getProduct)
		productAPI.POST("/product", createProduct)
	}
}

// Llis - list api
func getProduct(c *gin.Context) {
	c.JSON(200, gin.H{"result": "get product"})
}

// Create - create api
func createProduct(c *gin.Context) {
	c.JSON(200, gin.H{"result": "create product"})
}
